package usecase

import (
	"context"
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	producer "project/internal/qaas/send_messages/producer/usecase"
)

type usecase struct {
	producer     *amqp.Connection
	channel      *amqp.Channel
	queue        *amqp.Queue
	messagesRepo messages.Repository
}

func NewProducer(connAddr string, queueName string, messagesRepo messages.Repository) (producer.Usecase, error) {
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

	return &usecase{producer: producer, channel: channel, queue: &queue, messagesRepo: messagesRepo}, nil
}

func (u *usecase) ProduceMessage(ctx context.Context, producerMessage model.ProducerMessage) error {
	switch producerMessage.Type {
	case configs.Create:
		go func() {
			_, err := u.messagesRepo.InsertMessageInDB(ctx, model.Message{
				Id:        producerMessage.Id,
				Body:      producerMessage.Body,
				AuthorId:  producerMessage.AuthorId,
				ChatId:    producerMessage.ChatID,
				CreatedAt: producerMessage.CreatedAt,
			})
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Edit:
		go func() {
			_, err := u.messagesRepo.EditMessageById(ctx, producerMessage)
			if err != nil {
				log.Error(err)
			}
		}()
	case configs.Delete:
		go func() {
			err := u.messagesRepo.DeleteMessageById(ctx, producerMessage.Id)
			if err != nil {
				log.Error(err)
			}
		}()
	default:
		return errors.New("не выбран ни один из трех 0, 1, 2")
	}

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
