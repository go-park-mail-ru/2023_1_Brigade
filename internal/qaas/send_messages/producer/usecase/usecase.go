package usecase

import (
	"context"
	"os"
	"os/signal"
	"project/internal/qaas/send_messages/producer"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type usecase struct {
	producer sarama.AsyncProducer
	closed   bool
}

func NewProducer(brokerList []string) (producer.Usecase, error) {
	config := sarama.NewConfig()                     // Создаем конфигурацию для Kafka-продюсера
	config.Producer.Retry.Max = 5                    // Устанавливаем максимальное количество попыток ретрая
	config.Producer.RequiredAcks = sarama.WaitForAll // Устанавливаем ожидание подтверждения от всех брокеров кластера
	config.Producer.Return.Successes = true          // Включаем возвращение успешных сообщений в канал
	config.Producer.Return.Errors = true             // Включаем возвращение ошибок в канал

	time.Sleep(10 * time.Second)

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return &usecase{producer: nil}, err
	}

	return &usecase{producer: producer}, nil
}

func (u *usecase) ProduceMessage(ctx context.Context, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: "messages",
		Value: sarama.ByteEncoder(message),
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for {
			select {
			case err := <-u.producer.Errors():
				log.Error(err)
				u.producer.Input() <- msg
			case <-signals:
				if !u.closed {
					u.closed = true
					u.producer.AsyncClose()
				}
				log.Fatal()
			}
		}
	}()

	u.producer.Input() <- msg
	_ = <-u.producer.Successes()

	return nil
}
