package usecase

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/messages"
	"project/internal/model"
	"project/internal/qaas/send_messages/consumer"
	"project/internal/qaas/send_messages/producer"
)

type usecase struct {
	repo     messages.Repository
	producer producer.Usecase
	consumer consumer.Usecase
}

func NewMessagesUsecase(messagesRepo messages.Repository) messages.Usecase {
	brokerList := []string{"localhost:9092"}
	groupID := "group-message"

	producer, err := producer.NewProducer(brokerList)
	if err != nil {
		log.Fatal("producer:  ", err)
	}

	consumer, err := consumer.NewConsumer(brokerList, groupID)
	if err != nil {
		log.Fatal("consumer:  ", err)
	}

	consumer.StartConsumeMessages()

	return &usecase{repo: messagesRepo, producer: producer, consumer: consumer}
}

func (u *usecase) SendMessage(ctx echo.Context, jsonWebSocketMessage []byte) error {
	var webSocketMessage model.WebSocketMessage
	err := json.Unmarshal(jsonWebSocketMessage, &webSocketMessage)
	if err != nil {
		return err
	}

	chat, err := u.repo.GetChatById(webSocketMessage.ChatID)
	if err != nil {
		return err
	}

	message := model.Message{
		Id:       0,
		Body:     webSocketMessage.Body,
		AuthorId: webSocketMessage.AuthorID,
		ChatId:   webSocketMessage.ChatID,
		IsRead:   false,
	}
	message, err = u.repo.InsertMessageInDB(message)
	if err != nil {
		return err
	}

	for _, member := range chat.Members {
		producerMessage := model.ProducerMessage{
			Message:    message,
			ReceiverID: member.Id,
		}
		jsonProducerMessage, err := json.Marshal(producerMessage)
		if err != nil {
			return err
		}

		err = u.producer.ProduceMessage(jsonProducerMessage)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *usecase) ReceiveMessage(ctx echo.Context) ([]byte, error) {
	var message model.ProducerMessage
	jsonMessage := u.consumer.ConsumeMessage()

	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		return nil, err
	}

	err = u.repo.InsertMessageReceiveInDB(message)
	if err != nil {
		return nil, err
	}

	err = u.repo.MarkMessageReading(message.Message.Id)
	return jsonMessage, err
}
