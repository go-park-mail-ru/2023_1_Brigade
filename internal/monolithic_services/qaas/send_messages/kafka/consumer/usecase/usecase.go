package usecase

import (
	"context"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	consumer "project/internal/microservices/consumer/usecase"
)

type usecase struct {
	consumer     sarama.ConsumerGroup
	messagesChan chan []byte
}

type messageHandler struct {
	messagesChan chan []byte
}

func (h *messageHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *messageHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *messageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		session.MarkMessage(msg, "")
		h.messagesChan <- msg.Value
	}
	return nil
}

func NewConsumer(brokerList []string, groupID string) (consumer.Usecase, error) {
	messagesChan := make(chan []byte)

	config := sarama.NewConfig()                          // Создаем конфигурацию для Kafka-продюсера
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Начинаем с самого старого сообщения
	config.Consumer.Return.Errors = true                  // Включаем отслеживание ошибок

	consumer, err := sarama.NewConsumerGroup(brokerList, groupID, config)
	if err != nil {
		log.Fatal(err)
	}

	consumerUsecase := usecase{consumer: consumer, messagesChan: messagesChan}

	consumerUsecase.StartConsumeMessages(context.TODO())

	return &consumerUsecase, nil
}

func (u *usecase) ConsumeMessage(ctx context.Context) []byte {
	msg := <-u.messagesChan
	return msg
}

func (u *usecase) StartConsumeMessages(ctx context.Context) {
	handler := messageHandler{messagesChan: u.messagesChan}
	topic := []string{"messages"}

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		err := u.consumer.Close()
		if err != nil {
			log.Error(err)
		}
		log.Fatal()
	}()

	go func() {
		for err := range u.consumer.Errors() {
			log.Error(err)
		}
	}()

	go func() {
		err := u.consumer.Consume(ctx, topic, &handler)
		if err != nil {
			log.Error(err)
		}
	}()
}
