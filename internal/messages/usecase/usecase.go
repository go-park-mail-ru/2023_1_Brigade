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

func NewMessagesUsecase(chatRepo chat.Repository, messagesRepo messages.Repository, config configs.Kafka) messages.Usecase {
	consumer, err := consumerUsecase.NewConsumer(config.BrokerList, config.GroupID)
	if err != nil {
		log.Error(err)
	}

	producer, err := producerUsecase.NewProducer(config.BrokerList)
	if err != nil {
		log.Error(err)
	}

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

	consumer.StartConsumeMessages()

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
	//
	//var webSocketMessage model.WebSocketMessage
	//err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	//if err != nil {
	//	return err
	//}
	//
	//id := webSocketMessage.Id
	////createdAt := time.Now()
	//
	//// если пришел ивент на создание сообщения (0)
	//if id == "" {
	//	id = uuid.New().String()
	//}
	//
	//producerMessage := model.ProducerMessage{
	//	Id:       id,
	//	Type:     webSocketMessage.Type,
	//	Body:     webSocketMessage.Body,
	//	AuthorId: webSocketMessage.AuthorID,
	//	ChatID:   webSocketMessage.ChatID,
	//	//CreatedAt: createdAt,
	//}
	//
	//switch producerMessage.Type {
	//case configs.Create:
	//	go func() {
	//		_, err = u.messagesRepo.InsertMessageInDB(ctx, model.Message{
	//			Id:       id,
	//			Body:     producerMessage.Body,
	//			AuthorId: producerMessage.AuthorId,
	//			ChatId:   producerMessage.ChatID,
	//			//CreatedAt: createdAt,
	//		})
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}()
	//case configs.Edit:
	//	go func() {
	//		_, err = u.messagesRepo.EditMessageById(ctx, producerMessage)
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}()
	//case configs.Delete:
	//	go func() {
	//		err = u.messagesRepo.DeleteMessageById(ctx, id)
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}()
	//default:
	//	return errors.New("не выбран ни один из трех 0, 1, 2")
	//}
	//
	//return u.PutInProducer(ctx, producerMessage)

	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	id := webSocketMessage.Id
	createdAt := time.Now()

	if id == "" {
		id = uuid.New().String()
	}

	switch webSocketMessage.Type {
	case configs.Create:
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

		return u.PutInProducer(ctx, model.ProducerMessage{
			Id:        webSocketMessage.Id,
			Type:      webSocketMessage.Type,
			Body:      webSocketMessage.Body,
			AuthorId:  webSocketMessage.AuthorID,
			ChatID:    webSocketMessage.ChatID,
			CreatedAt: createdAt,
		})
	}

	return errors.New("не выбран ни один из трех 0, 1, 2")
}

func (u usecase) PutInProducer(ctx context.Context, producerMessage model.ProducerMessage) error {
	//members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), producerMessage.ChatID)
	//if err != nil {
	//	return err
	//}
	//
	//for _, member := range members {
	//	if member.MemberId == producerMessage.AuthorId {
	//		continue
	//	}
	//
	//	jsonProducerMessage, err := json.Marshal(producerMessage)
	//	if err != nil {
	//		return err
	//	}
	//
	//	err = u.producer.ProduceMessage(jsonProducerMessage)
	//	if err != nil {
	//		return err
	//	}
	//
	//	err = u.centrifugePublication(jsonProducerMessage)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//return nil
	members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), producerMessage.ChatID)
	if err != nil {
		return err
	}

	//id := uuid.New().String()
	//createdAt := time.Now()

	//go func() {
	//	message := model.Message{
	//		Id:        id,
	//		Body:      producerMessage.Body,
	//		AuthorId:  producerMessage.AuthorId,
	//		ChatId:    producerMessage.ChatID,
	//		CreatedAt: createdAt,
	//	}
	//
	//	_, err = u.messagesRepo.InsertMessageInDB(context.Background(), message)
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}()

	for _, member := range members {
		if member.MemberId == producerMessage.AuthorId {
			continue
		}

		//producerMessage := model.ProducerMessage{
		//	Id:         producerMessage.Id,
		//	Body:       producerMessage.Body,
		//	AuthorId:   producerMessage.AuthorId,
		//	ChatID:     producerMessage.ChatID,
		//	ReceiverID: member.MemberId,
		//	CreatedAt:  producerMessage.CreatedAt,
		//}
		producerMessage.ReceiverID = member.MemberId
		jsonProducerMessage, err := json.Marshal(producerMessage)
		if err != nil {
			return err
		}

		err = u.producer.ProduceMessage(jsonProducerMessage)
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
	jsonMessage := u.consumer.ConsumeMessage()

	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		return nil, err
	}

	return jsonMessage, nil
}
