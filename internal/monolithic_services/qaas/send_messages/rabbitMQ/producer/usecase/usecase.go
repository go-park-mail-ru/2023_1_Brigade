package usecase

import (
	"context"
	"os"
	"os/signal"
	producer "project/internal/microservices/producer/usecase"
	"project/internal/model"

	"github.com/mailru/easyjson"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type usecase struct {
	producer *amqp.Connection
	channel  *amqp.Channel
	queue    *amqp.Queue
	dlxQueue *amqp.Queue
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

	//err = channel.ExchangeDeclare(
	//	"dlx_exchange",
	//	"fanout",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil)
	//
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}
	//
	//dlxQueue, err := channel.QueueDeclare(
	//	"dlx_queue",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}

	//err = channel.QueueBind(
	//	"dlx_queue",
	//	"dlx-routing-key",
	//	"dlx_exchange",
	//	false,
	//	nil,
	//)

	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}

	err = channel.ExchangeDeclare(
		"messages_exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
		//amqp.Table{
		//	"x-dead-letter-exchange":    "dlx_exchange",
		//	"x-dead-letter-routing-key": "dlx-routing-key",
		//},
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = channel.QueueBind(
		queueName,
		"",
		"messages_exchange",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
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

	return &usecase{producer: producer, channel: channel, queue: &queue, dlxQueue: nil}, nil
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
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
