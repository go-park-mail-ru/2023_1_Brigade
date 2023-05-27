package usecase

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	consumer "project/internal/microservices/consumer/usecase"
	"project/internal/monolithic_services/centrifugo"
)

type usecase struct {
	consumer *amqp.Connection
	channel  *amqp.Channel
	//queue       *amqp.Queue
	client      centrifugo.Centrifugo
	channelName string
}

func NewConsumer(connAddr string, queueName string, centrifugo centrifugo.Centrifugo, channelName string) (consumer.Usecase, error) {
	consumer, err := amqp.Dial(connAddr)
	if err != nil {
		return nil, err
	}

	channel, err := consumer.Channel()
	if err != nil {
		return nil, err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		err = consumer.Close()
		if err != nil {
			log.Error(err)
		}

		err = channel.Close()
		if err != nil {
			log.Error(err)
		}

		centrifugo.Close()
		log.Fatal()
	}()

	consumerUsecase := usecase{consumer: consumer, channel: channel, client: centrifugo, channelName: channelName}

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
			"messages",
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
			err = u.centrifugePublication(msg.Body)
			if err != nil {
				log.Error(err)
			}

			if err != nil {
				log.Error(err)
			}
		}
	}
}
