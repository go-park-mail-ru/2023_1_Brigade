package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/microservices/chat"
	consumer "project/internal/microservices/consumer/usecase"
	"project/internal/microservices/messages"
	producer "project/internal/microservices/producer/usecase"
	"project/internal/model"
	httpUtils "project/internal/pkg/http_utils"
	"time"
)

type usecase struct {
	chatRepo     chat.Repository
	messagesRepo messages.Repository
	producer     producer.Usecase
	consumer     consumer.Usecase
}

func NewMessagesUsecase(chatRepo chat.Repository, consumer consumer.Usecase, producer producer.Usecase, messagesRepo messages.Repository) messages.Usecase {
	signals := make(chan os.Signal, 1)
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

	webSocketMessage = httpUtils.SanitizeStruct(webSocketMessage).(model.WebSocketMessage)

	producerMessage := model.ProducerMessage{
		Id:          webSocketMessage.Id,
		Attachments: webSocketMessage.Attachments,
		Action:      webSocketMessage.Action,
		Type:        webSocketMessage.Type,
		Body:        webSocketMessage.Body,
		AuthorId:    webSocketMessage.AuthorID,
		ChatID:      webSocketMessage.ChatID,
	}

	// если пришел ивент на создание сообщения (0)
	if producerMessage.Action == config.Create {
		producerMessage.Id = uuid.NewString()
		producerMessage.CreatedAt = time.Now().String()
	}

	switch producerMessage.Action {
	case config.Create:
		go func() {
			err := u.messagesRepo.InsertMessageInDB(context.TODO(), model.Message{
				Id:          producerMessage.Id,
				Attachments: webSocketMessage.Attachments,
				Type:        producerMessage.Type,
				Body:        producerMessage.Body,
				AuthorId:    producerMessage.AuthorId,
				ChatId:      producerMessage.ChatID,
				CreatedAt:   producerMessage.CreatedAt,
			})
			if err != nil {
				log.Error(err)
			}
		}()
	case config.Edit:
		go func() {
			_, err := u.messagesRepo.EditMessageById(context.TODO(), producerMessage)
			if err != nil {
				log.Error(err)
			}
		}()
	case config.Delete:
		go func() {
			err := u.messagesRepo.DeleteMessageById(context.TODO(), producerMessage.Id)
			if err != nil {
				log.Error(err)
			}
		}()
	default:
		return errors.New("не выбран ни один из трех 0, 1, 2")
	}

	members, err := u.chatRepo.GetChatMembersByChatId(ctx, webSocketMessage.ChatID)
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
