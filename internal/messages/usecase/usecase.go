package usecase

import (
	"context"
	"encoding/json"
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

	return &usecase{chatRepo: chatRepo, messagesRepo: messagesRepo, producer: producer, consumer: consumer}
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
	}

	config := centrifuge.Config{}
	c := centrifuge.NewJsonClient("ws://centrifugo:8900/connection/websocket", config) // url - адрес на котором работает centrifugo
	c.OnError(func(e centrifuge.ErrorEvent) {
		log.Error(e.Error)
	})
	c.OnConnecting(func(e centrifuge.ConnectingEvent) {
		log.Warn("Connecting - %d (%s)", e.Code, e.Reason)
	})
	c.OnConnected(func(e centrifuge.ConnectedEvent) {
		log.Warn("Connected with ID %s", e.ClientID)
	})
	c.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		log.Warn("Disconnected: %d (%s)", e.Code, e.Reason)
	})

	c.OnError(func(e centrifuge.ErrorEvent) {
		log.Error("Error: %s", e.Error.Error())
	})

	c.OnMessage(func(e centrifuge.MessageEvent) {
		log.Printf("Message from server: %s", string(e.Data))
	})

	c.OnSubscribed(func(e centrifuge.ServerSubscribedEvent) {
		log.Warn("Subscribed to server-side channel %s: (was recovering: %v, recovered: %v)", e.Channel, e.WasRecovering, e.Recovered)
	})
	c.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
		log.Warn("Subscribing to server-side channel %s", e.Channel)
	})
	c.OnUnsubscribed(func(e centrifuge.ServerUnsubscribedEvent) {
		log.Warn("Unsubscribed from server-side channel %s", e.Channel)
	})

	c.OnPublication(func(e centrifuge.ServerPublicationEvent) {
		log.Warn("Publication from server-side channel %s: %s (offset %d)", e.Channel, e.Data, e.Offset)
	})
	c.OnJoin(func(e centrifuge.ServerJoinEvent) {
		log.Warn("Join to server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
	})
	c.OnLeave(func(e centrifuge.ServerLeaveEvent) {
		log.Warn("Leave from server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
	})
	defer c.Close()

	err := c.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := c.NewSubscription("channel", centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	sub.OnSubscribing(func(e centrifuge.SubscribingEvent) {
		log.Warn("Subscribing on channel %s - %d (%s)", sub.Channel, e.Code, e.Reason)
	})
	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		log.Warn("Subscribed on channel %s, (was recovering: %v, recovered: %v)", sub.Channel, e.WasRecovering, e.Recovered)
	})
	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		log.Warn("Unsubscribed from channel %s - %d (%s)", sub.Channel, e.Code, e.Reason)
	})

	sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
		log.Error("Subscription error %s: %s", sub.Channel, e.Error)
	})

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Warn("опубликовал")
	})

	sub.OnJoin(func(e centrifuge.JoinEvent) {
		log.Warn("Someone joined %s: user id %s, client id %s", sub.Channel, e.User, e.Client)
	})
	sub.OnLeave(func(e centrifuge.LeaveEvent) {
		log.Warn("Someone left %s: user id %s, client id %s", sub.Channel, e.User, e.Client)
	})

	err = sub.Subscribe()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := sub.Publish(context.Background(), []byte{})
	log.Error(res, err)

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
