package usecase

import (
	"context"
	"encoding/json"
	"errors"
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
}

func NewMessagesUsecase(chatRepo chat.Repository, consumer consumer.Usecase, producer producer.Usecase, messagesRepo messages.Repository) messages.Usecase {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		log.Fatal()
	}()

	return &usecase{chatRepo: chatRepo, messagesRepo: messagesRepo, producer: producer, consumer: consumer}
}

func (u usecase) PutInProducer(ctx context.Context, jsonWebSocketMessage []byte) error {
	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	producerMessage := model.ProducerMessage{
		Id:        webSocketMessage.Id,
		Type:      webSocketMessage.Type,
		Body:      webSocketMessage.Body,
		AuthorId:  webSocketMessage.AuthorID,
		ChatID:    webSocketMessage.ChatID,
		CreatedAt: time.Now().String(),
	}

	// если пришел ивент на создание сообщения (0)
	if producerMessage.Id == "" {
		producerMessage.Id = uuid.New().String()
	}

	switch producerMessage.Type {
	case configs.Create:
		go func() {
			err := u.messagesRepo.InsertMessageInDB(context.TODO(), model.Message{
				Id:        producerMessage.Id,
				Body:      producerMessage.Body,
				AuthorId:  producerMessage.AuthorId,
				ChatId:    producerMessage.ChatID,
				CreatedAt: producerMessage.CreatedAt,
			})
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Edit:
		go func() {
			_, err := u.messagesRepo.EditMessageById(context.TODO(), producerMessage)
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Delete:
		go func() {
			err := u.messagesRepo.DeleteMessageById(context.TODO(), producerMessage.Id)
			if err != nil {
				log.Error(err)
			}
		}()
	default:
		return errors.New("не выбран ни один из трех 0, 1, 2")
	}

	members, err := u.chatRepo.GetChatMembersByChatId(context.Background(), webSocketMessage.ChatID)
	if err != nil {
		return err
	}

	for _, member := range members {
		producerMessage.ReceiverID = member.MemberId
		err = u.producer.ProduceMessage(ctx, producerMessage)
		if err != nil {
			return err
		}
	}

	return nil
}
