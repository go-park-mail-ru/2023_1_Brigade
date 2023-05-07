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
	consumer "project/internal/qaas/send_messages/consumer/usecase"
	producer "project/internal/qaas/send_messages/producer/usecase"
	"time"
)

type usecase struct {
	chatRepo     chat.Repository
	messagesRepo messages.Repository
	producer     producer.Usecase
	consumer     consumer.Usecase
	client       *centrifuge.Client
}

func NewMessagesUsecase(chatRepo chat.Repository, messagesRepo messages.Repository, consumer consumer.Usecase, producer producer.Usecase) messages.Usecase {
	c := centrifuge.NewJsonClient("ws://localhost:8900/connection/websocket", centrifuge.Config{})
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		_ = c.Close
		log.Fatal()
	}()

	err := c.Connect()
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

func (u usecase) SwitchMessageType(ctx context.Context, jsonWebSocketMessage []byte) error {
	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	id := webSocketMessage.Id
	createdAt := time.Now()

	// если пришел ивент на создание сообщения (0)
	if id == "" {
		id = uuid.New().String()
	}

	producerMessage := model.ProducerMessage{
		Id:       id,
		Type:     webSocketMessage.Type,
		Body:     webSocketMessage.Body,
		AuthorId: webSocketMessage.AuthorID,
		ChatID:   webSocketMessage.ChatID,
	}

	if id == "" {
		producerMessage.CreatedAt = createdAt
	}

	switch producerMessage.Type {
	case configs.Create:
		go func() {
			_, err = u.messagesRepo.InsertMessageInDB(ctx, model.Message{
				Id:        id,
				Body:      producerMessage.Body,
				AuthorId:  producerMessage.AuthorId,
				ChatId:    producerMessage.ChatID,
				CreatedAt: createdAt,
			})
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Edit:
		go func() {
			_, err = u.messagesRepo.EditMessageById(ctx, producerMessage)
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Delete:
		go func() {
			err = u.messagesRepo.DeleteMessageById(ctx, id)
			if err != nil {
				log.Error(err)
			}
		}()
	default:
		return errors.New("не выбран ни один из трех 0, 1, 2")
	}

	return u.PutInProducer(ctx, producerMessage)
}

func (u usecase) PutInProducer(ctx context.Context, producerMessage model.ProducerMessage) error {
	members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), producerMessage.ChatID)
	if err != nil {
		return err
	}

	for _, member := range members {
		//if member.MemberId == producerMessage.AuthorId {
		//	continue
		//}

		producerMessage.ReceiverID = member.MemberId
		jsonProducerMessage, err := json.Marshal(producerMessage)
		if err != nil {
			return err
		}

		err = u.producer.ProduceMessage(ctx, jsonProducerMessage)
		if err != nil {
			return err
		}

		jsonWebSocketMessage, err := json.Marshal(producerMessage)
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

func (u usecase) PullFromConsumer(ctx context.Context) ([]byte, error) {
	var message model.ProducerMessage
	jsonMessage := u.consumer.ConsumeMessage(ctx)

	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		return nil, err
	}

	return jsonMessage, nil
}
