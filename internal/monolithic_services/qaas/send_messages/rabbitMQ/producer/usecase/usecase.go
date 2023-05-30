package usecase

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	producer "project/internal/microservices/producer/usecase"
	"project/internal/model"

	"github.com/mailru/easyjson"

	amqp "github.com/rabbitmq/amqp091-go"
)

type usecase struct {
	channel *amqp.Channel
	queue   *amqp.Queue
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

	err = channel.ExchangeDeclare(
		"user_dlx",
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		"user_create_dlx",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind(
		"user_create_dlx",
		"",
		"user_dlx",
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		amqp.Table{"x-dead-letter-exchange": "user_dlx"},
	)
	if err != nil {
		return nil, err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		err = producer.Close()
		if err != nil {
			log.Error(err)
		}

		err = channel.Close()
		if err != nil {
			log.Error(err)
		}

		log.Fatal()
	}()

	return &usecase{channel: channel, queue: &queue}, nil
}

func (u *usecase) ProduceMessage(ctx context.Context, producerMessage model.ProducerMessage) error {
	message, err := easyjson.Marshal(producerMessage)
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
			ContentType: "application/json",
			Body:        message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
