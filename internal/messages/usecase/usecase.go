package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/centrifugal/centrifuge-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	"project/internal/qaas/send_messages/consumer"
	consumerUsecase "project/internal/qaas/send_messages/consumer/usecase"
	"project/internal/qaas/send_messages/producer"
	producerUsecase "project/internal/qaas/send_messages/producer/usecase"
)

type usecase struct {
	chatRepo     chat.Repository
	messagesRepo messages.Repository
	producer     producer.Usecase
	consumer     consumer.Usecase
	client       *centrifuge.Client
}

func NewMessagesUsecase(chatRepo chat.Repository, messagesRepo messages.Repository, config configs.Kafka) messages.Usecase {
	consumer, err := consumerUsecase.NewConsumer(config.BrokerList, config.GroupID)
	if err != nil {
		log.Error(err)
	}

	producer, err := producerUsecase.NewProducer(config.BrokerList)
	if err != nil {
		log.Error(err)
	}

	consumer.StartConsumeMessages()

	c := centrifuge.NewJsonClient("ws://centrifugo:8900/connection/websocket", centrifuge.Config{})
	defer c.Close()

	err = c.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := c.NewSubscription("channel", centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Error(err)
	}

	err = sub.Subscribe()
	if err != nil {
		log.Error(err)
	}

	return &usecase{chatRepo: chatRepo, messagesRepo: messagesRepo, producer: producer, consumer: consumer, client: c}
}

func (u usecase) centrifugePublication(jsonWebSocketMessage []byte) error {
	sub, subscribed := u.client.GetSubscription("channel")
	if !subscribed {
		return errors.New("не подписан")
	}

	_, err := sub.Publish(context.Background(), jsonWebSocketMessage)
	return err
}

func (u usecase) SendMessage(ctx echo.Context, jsonWebSocketMessage []byte) error {
	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), webSocketMessage.ChatID)
	if err != nil {
		return err
	}

	go func() {
		message := model.Message{
			Id:       0,
			Body:     webSocketMessage.Body,
			AuthorId: webSocketMessage.AuthorID,
			ChatId:   webSocketMessage.ChatID,
		}

		_, err = u.messagesRepo.InsertMessageInDB(context.Background(), message)
		if err != nil {
			log.Error(err)
		}
	}()

	for _, member := range members {
		if member.MemberId == webSocketMessage.AuthorID {
			continue
		}

		producerMessage := model.ProducerMessage{
			Body:       webSocketMessage.Body,
			AuthorId:   webSocketMessage.AuthorID,
			ChatID:     webSocketMessage.ChatID,
			ReceiverID: member.MemberId,
		}
		jsonProducerMessage, err := json.Marshal(producerMessage)
		if err != nil {
			return err
		}

		err = u.producer.ProduceMessage(jsonProducerMessage)
		if err != nil {
			return err
		}

		err = u.centrifugePublication(jsonWebSocketMessage)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u usecase) ReceiveMessage(ctx echo.Context) ([]byte, error) {
	var message model.ProducerMessage
	jsonMessage := u.consumer.ConsumeMessage()

	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		return nil, err
	}

	return jsonMessage, nil
}
