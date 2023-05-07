package usecase

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/chat"
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

func NewMessagesUsecase(chatRepo chat.Repository, messagesRepo messages.Repository, consumer consumer.Usecase, producer producer.Usecase) messages.Usecase {
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
