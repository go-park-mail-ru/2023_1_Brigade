package usecase

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	producer "project/internal/microservices/producer/usecase"
	"project/internal/model"
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
			producer.Close()
			channel.Close()
			log.Fatal()
		}
	}()

	return &usecase{producer: producer, channel: channel, queue: &queue}, nil
}

func (u *usecase) ProduceMessage(ctx context.Context, producerMessage model.ProducerMessage) error {
	message, err := json.Marshal(producerMessage)
	if err != nil {
		return err
	}

	err = u.channel.PublishWithContext(
		ctx,
		"",
		u.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
