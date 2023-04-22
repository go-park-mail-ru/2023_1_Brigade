package usecase

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/centrifugal/centrifuge-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/qaas/send_messages/consumer"
)

type usecase struct {
	consumer     sarama.ConsumerGroup
	messagesChan chan []byte
}

type messageHandler struct {
	messagesChan chan []byte
}

func (h *messageHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *messageHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *messageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		session.MarkMessage(msg, "")
		h.messagesChan <- msg.Value
	}
	return nil
}

func NewConsumer(brokerList []string, groupID string) (consumer.Usecase, error) {
	messagesChan := make(chan []byte, 10)

	config := sarama.NewConfig()                          // Создаем конфигурацию для Kafka-продюсера
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Начинаем с самого старого сообщения
	config.Consumer.Return.Errors = true                  // Включаем отслеживание ошибок

	consumer, err := sarama.NewConsumerGroup(brokerList, groupID, config)
	if err != nil {
		return &usecase{}, err
	}

	return &usecase{consumer: consumer, messagesChan: messagesChan}, nil
}

func (u *usecase) ConsumeMessage() []byte {
	config := centrifuge.Config{}
	c := centrifuge.NewProtobufClient("centrifugo:8900", config) // url - адрес на котором работает centrifugo
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

	msg := <-u.messagesChan
	return msg
}

func (u *usecase) StartConsumeMessages() {
	handler := messageHandler{messagesChan: u.messagesChan}
	ctx := context.Background()
	topic := []string{"message"}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		err := u.consumer.Close()
		if err != nil {
			log.Error(err)
		}
		log.Fatal()
	}()

	go func() {
		for err := range u.consumer.Errors() {
			log.Error(err)
		}
	}()

	go func() {
		err := u.consumer.Consume(ctx, topic, &handler)
		if err != nil {
			log.Error(err)
		}
	}()
}
