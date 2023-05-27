//package usecase
//
//import (
//	"context"
//	"errors"
//	amqp "github.com/rabbitmq/amqp091-go"
//	log "github.com/sirupsen/logrus"
//	"os"
//	"os/signal"
//	consumer "project/internal/microservices/consumer/usecase"
//	"project/internal/monolithic_services/centrifugo"
//)
//
//type usecase struct {
//	consumer    *amqp.Connection
//	channel     *amqp.Channel
//	queue       *amqp.Queue
//	client      centrifugo.Centrifugo
//	channelName string
//}
//
//func NewConsumer(connAddr string, queueName string, centrifugo centrifugo.Centrifugo, channelName string) (consumer.Usecase, error) {
//	consumer, err := amqp.Dial(connAddr)
//	if err != nil {
//		return nil, err
//	}
//
//	channel, err := consumer.Channel()
//	if err != nil {
//		return nil, err
//	}
//
//	queue, err := channel.QueueDeclare(
//		queueName,
//		true,
//		false,
//		false,
//		true,
//		nil,
//		//amqp.Table{
//		//	"x-dead-letter-exchange":    "dlx_exchange",
//		//	"x-dead-letter-routing-key": "dlx-routing-key",
//		//},
//	)
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	go func() {
//		<-signals
//		err = consumer.Close()
//		if err != nil {
//			log.Error(err)
//		}
//
//		err = channel.Close()
//		if err != nil {
//			log.Error(err)
//		}
//
//		centrifugo.Close()
//		log.Fatal()
//	}()
//
//	consumerUsecase := usecase{consumer: consumer, channel: channel, queue: &queue, client: centrifugo, channelName: channelName}
//
//	go func() {
//		consumerUsecase.StartConsumeMessages(context.TODO())
//	}()
//
//	return &consumerUsecase, nil
//}
//
//func (u *usecase) centrifugePublication(jsonWebSocketMessage []byte) error {
//	sub, subscribed := u.client.GetSubscription(u.channelName)
//	if !subscribed {
//		return errors.New("не подписан")
//	}
//
//	_, err := sub.Publish(context.TODO(), jsonWebSocketMessage)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (u *usecase) StartConsumeMessages(ctx context.Context) {
//	for {
//		msgs, err := u.channel.Consume(
//			u.queue.Name,
//			"",
//			true,
//			false,
//			false,
//			true,
//			nil,
//		)
//
//		if err != nil {
//			log.Error(err)
//		}
//		//go func() {
//		for msg := range msgs {
//			err = u.centrifugePublication(msg.Body)
//			if err != nil {
//				log.Error(err)
//			}
//
//			if err != nil {
//				log.Error(err)
//			}
//
//			if err != nil {
//				log.Error(err)
//			}
//		}
//		//}()
//	}
//}

package usecase

import (
	"context"
	//"errors"
	//amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	amqp "github.com/wagslane/go-rabbitmq"
	"os"
	"os/signal"
	consumer "project/internal/microservices/consumer/usecase"
	"project/internal/monolithic_services/centrifugo"
)

type usecase struct {
	consumer *amqp.Consumer
	//channel     *amqp.Channel
	//queue       *amqp.Queue
	client      centrifugo.Centrifugo
	channelName string
}

func NewConsumer(connAddr string, queueName string, centrifugo centrifugo.Centrifugo, channelName string) (consumer.Usecase, error) {
	//consumer, err := amqp.Dial(connAddr)
	//if err != nil {
	//	return nil, err
	//}
	//
	//channel, err := consumer.Channel()
	//if err != nil {
	//	return nil, err
	//}
	//
	//queue, err := channel.QueueDeclare(
	//	queueName,
	//	false,
	//	false,
	//	false,
	//	true,
	//	nil,
	//)
	//if err != nil {
	//	return nil, err
	//}

	conn, err := rabbitmq.NewConn(connAddr, amqp.WithConnectionOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			//err := (msg.Body)
			//if err != nil {
			//	log.Error(err)
			//}
			//log.Printf("consumed: %v", string(d.Body))
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue

			sub, subscribed := centrifugo.GetSubscription(channelName)
			if !subscribed {
				log.Error("не подписан")
				//return errors.New("не подписан")
			}

			_, err := sub.Publish(context.TODO(), d.Body)
			if err != nil {
				log.Error(err)
				//return err
			}

			//return nil

			return rabbitmq.Ack
		},
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		consumer.Close()
		log.Fatal()
		//err = consumer.Close()
		//if err != nil {
		//	log.Error(err)
		//}
		//
		//err = channel.Close()
		//if err != nil {
		//	log.Error(err)
		//}
		//
		//centrifugo.Close()
		//log.Fatal()
	}()

	consumerUsecase := usecase{consumer: consumer, client: centrifugo, channelName: channelName}

	//go func() {
	//	consumerUsecase.StartConsumeMessages(context.TODO())
	//}()

	return &consumerUsecase, nil
}

//func (u *usecase) centrifugePublication(jsonWebSocketMessage []byte) error {
//	sub, subscribed := u.client.GetSubscription(u.channelName)
//	if !subscribed {
//		return errors.New("не подписан")
//	}
//
//	_, err := sub.Publish(context.TODO(), jsonWebSocketMessage)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (u *usecase) StartConsumeMessages(ctx context.Context) {
	//for {
	//	msgs, err := u.channel.Consume(
	//		u.queue.Name,
	//		"",
	//		true,
	//		false,
	//		false,
	//		true,
	//		nil,
	//	)
	//
	//	if err != nil {
	//		log.Error(err)
	//	}
	//
	//	for msg := range msgs {
	//		err := u.centrifugePublication(msg.Body)
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}
	//}
}
