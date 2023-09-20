package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/microservices/chat"
	"project/internal/microservices/messages"
	"project/internal/model"
	"project/internal/monolithic_services/centrifugo"
	httpUtils "project/internal/pkg/http_utils"
	"time"
)

type usecase struct {
	chatRepo     chat.Repository
	messagesRepo messages.Repository
	client       centrifugo.Centrifugo
	channelName  string
}

func NewMessagesUsecase(chatRepo chat.Repository, client centrifugo.Centrifugo, channelName string, messagesRepo messages.Repository) messages.Usecase {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		log.Fatal()
	}()

	return &usecase{chatRepo: chatRepo, messagesRepo: messagesRepo, client: client, channelName: channelName}
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

	members, err := u.chatRepo.GetChatMembersByChatId(context.TODO(), webSocketMessage.ChatID)
	if err != nil {
		return err
	}

	for _, member := range members {
		producerMessage.ReceiverID = member.MemberId

		data, err := easyjson.Marshal(producerMessage)
		if err != nil {
			log.Error(err)
		}

		err = u.centrifugePublication(data)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func (u *usecase) centrifugePublication(jsonWebSocketMessage []byte) error {
	sub, subscribed := u.client.GetSubscription(u.channelName)
	if !subscribed {
		return errors.New("не подписан")
	}

	_, err := sub.Publish(context.TODO(), jsonWebSocketMessage)
	if err != nil {
		return err
	}

	return nil
}
