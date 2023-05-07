package usecase

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	producer "project/internal/qaas/send_messages/producer/usecase"
)

type usecase struct {
	producer *amqp.Connection
	channel  *amqp.Channel
	queue    *amqp.Queue
}

func NewProducer(connAddr string, queueName string) (producer.Usecase, error) {
	producer, err := amqp.Dial(connAddr)
	if err != nil {
		return nil, err
	}

	channel, err := producer.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		true,
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
			producer.Close()
			channel.Close()
			log.Fatal()
		}
	}()

	return &usecase{producer: producer, channel: channel, queue: &queue}, nil
}

func (u *usecase) ProduceMessage(ctx context.Context, message []byte) error {
	err := u.channel.PublishWithContext(
		ctx,
		"",
		u.queue.Name,
		true,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	log.Printf(" [x] Sent %s", message)
	if err != nil {
		log.Error("Failed to publish a message: %v", err)
		return err
	}

	return nil
}
