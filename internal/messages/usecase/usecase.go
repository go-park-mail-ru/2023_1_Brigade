package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/centrifugal/centrifuge-go"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	"project/internal/qaas/send_messages/consumer"
	consumerUsecase "project/internal/qaas/send_messages/consumer/usecase"
	"project/internal/qaas/send_messages/producer"
	producerUsecase "project/internal/qaas/send_messages/producer/usecase"
	"time"
)

type usecase struct {
	chatRepo     chat.Repository
	messagesRepo messages.Repository
	producer     producer.Usecase
	consumer     consumer.Usecase
	client       *centrifuge.Client
}

func (u usecase) EditMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	//TODO implement me
	panic("implement me")
}

func (u usecase) DeleteMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	//TODO implement me
	panic("implement me")
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
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		_ = c.Close
		log.Fatal()
	}()

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

func (u usecase) SwitchMesssageType(ctx context.Context, jsonWebSocketMessage []byte) error {
	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	switch webSocketMessage.Type {
	case configs.Create:
		return u.SendMessage(ctx, webSocketMessage)
	case configs.Edit:
		return u.EditMessage(ctx, webSocketMessage)
	case configs.Delete:
		return u.DeleteMessage(ctx, webSocketMessage)
	}

	return errors.New("не выбран ни один из трех 0, 1, 2")
}

func (u usecase) SendMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), webSocketMessage.ChatID)
	if err != nil {
		return err
	}

	id := uuid.New().String()
	createdAt := time.Now()

	go func() {
		message := model.Message{
			Id:        id,
			Body:      webSocketMessage.Body,
			AuthorId:  webSocketMessage.AuthorID,
			ChatId:    webSocketMessage.ChatID,
			CreatedAt: createdAt,
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
			Id:         id,
			Body:       webSocketMessage.Body,
			AuthorId:   webSocketMessage.AuthorID,
			ChatID:     webSocketMessage.ChatID,
			ReceiverID: member.MemberId,
			CreatedAt:  createdAt,
		}
		jsonProducerMessage, err := json.Marshal(producerMessage)
		if err != nil {
			return err
		}

		err = u.producer.ProduceMessage(jsonProducerMessage)
		if err != nil {
			return err
		}

		jsonWebSocketMessage, err := json.Marshal(webSocketMessage)
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

func (u usecase) ReceiveMessage(ctx context.Context) ([]byte, error) {
	var message model.ProducerMessage
	jsonMessage := u.consumer.ConsumeMessage()

	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		return nil, err
	}

	return jsonMessage, nil
}
