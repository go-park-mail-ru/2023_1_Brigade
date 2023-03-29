package usecase

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	"project/internal/qaas/send_messages/consumer"
	"project/internal/qaas/send_messages/producer"
	"time"
)

type usecase struct {
	repo     messages.Repository
	producer producer.Usecase
	consumer consumer.Usecase
}

func NewMessagesUsecase(messagesRepo messages.Repository, config configs.Kafka) messages.Usecase {
	//brokerList := []string{"host.docker.internal:9092"} //localhost - локально
	//groupID := "group-message"

	time.Sleep(time.Second * 1) // не успевает инициализироваться kafka в docker

	consumer, err := consumer.NewConsumer(config.BrokerList, config.GroupID)
	if err != nil {
		log.Error("consumer:  ", err)
	}

	producer, err := producer.NewProducer(config.BrokerList)
	if err != nil {
		log.Error("producer:  ", err) // log.Fatal
	}

	//consumer.StartConsumeMessages()

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

	go func() {
		message := model.Message{
			Id:       0,
			Body:     webSocketMessage.Body,
			AuthorId: webSocketMessage.AuthorID,
			ChatId:   webSocketMessage.ChatID,
			IsRead:   false,
		}

		_, err = u.repo.InsertMessageInDB(message)
		if err != nil {
			log.Error(err)
		}
	}()

	for _, member := range chat {
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

	go func() {
		err = u.repo.InsertMessageReceiveInDB(message)
		if err != nil {
			log.Error(err)
		}

		//err = u.repo.MarkMessageReading(message.)
		//if err != nil {
		//	log.Error(err)
		//}
	}()

	return jsonMessage, nil
}
