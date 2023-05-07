package usecase

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	consumer "project/internal/qaas/send_messages/consumer/usecase"
)

type usecase struct {
	consumer     *amqp.Connection
	channel      *amqp.Channel
	queue        *amqp.Queue
	messagesChan chan []byte
}

func NewConsumer(connAddr string, queueName string) (consumer.Usecase, error) {
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
			consumer.Close()
			channel.Close()
		}
	}()

	return &usecase{consumer: consumer, channel: channel, queue: &queue}, nil
}

func (u *usecase) ConsumeMessage(ctx context.Context) []byte {
	msg := <-u.messagesChan
	return msg
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
			u.messagesChan <- msg.Body
		}
	}
}
