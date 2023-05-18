package usecase

import (
	"context"
	"errors"
	"github.com/centrifugal/centrifuge-go"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/config"
	consumer "project/internal/microservices/consumer/usecase"
)

type usecase struct {
	consumer     *amqp.Connection
	channel      *amqp.Channel
	queue        *amqp.Queue
	messagesChan chan []byte
	client       *centrifuge.Client
	channelName  string
}

func NewConsumer(connAddr string, queueName string, centrifugo config.Centrifugo) (consumer.Usecase, error) {
	c := centrifuge.NewJsonClient(centrifugo.ConnAddr, centrifuge.Config{})

	err := c.Connect()
	if err != nil {
		log.Error(err)
	}

	sub, err := c.NewSubscription(centrifugo.ChannelName, centrifuge.SubscriptionConfig{
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

	consumer, err := amqp.Dial(connAddr)
	if err != nil {
		return nil, err
	}

	channel, err := consumer.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		select {
		case <-signals:
			consumer.Close()
			channel.Close()
			c.Close()
			log.Fatal()
		}
	}()

	consumerUsecase := usecase{consumer: consumer, channel: channel, queue: &queue, client: c, channelName: centrifugo.ChannelName}

	go func() {
		consumerUsecase.StartConsumeMessages(context.TODO())
	}()

	return &consumerUsecase, nil
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

func (u *usecase) StartConsumeMessages(ctx context.Context) {
	for {
		msgs, err := u.channel.Consume(
			u.queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Error(err)
		}

		for msg := range msgs {
			err := u.centrifugePublication(msg.Body)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
