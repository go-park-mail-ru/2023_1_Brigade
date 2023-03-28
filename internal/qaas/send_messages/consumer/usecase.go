package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

type Usecase struct {
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
		log.Printf("Message claimed: value = %s, topic = %s, partition = %d, offset = %d\n",
			string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
		session.MarkMessage(msg, "")
		h.messagesChan <- msg.Value
	}
	return nil
}

func NewConsumer(brokerList []string, groupID string) (Usecase, error) {
	messagesChan := make(chan []byte, 10)

	config := sarama.NewConfig()                          // Создаем конфигурацию для Kafka-продюсера
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Начинаем с самого старого сообщения
	config.Consumer.Return.Errors = true                  // Включаем отслеживание ошибок

	consumer, err := sarama.NewConsumerGroup(brokerList, groupID, config)
	if err != nil {
		log.Info("Failed to create consumer group: ", err)
		return Usecase{consumer: consumer, messagesChan: messagesChan}, err
	}

	return Usecase{consumer: consumer, messagesChan: messagesChan}, nil
}

func (u *Usecase) ConsumeMessage() []byte {
	msg := <-u.messagesChan
	return msg
}

func (u *Usecase) StartConsumeMessages() {
	handler := messageHandler{messagesChan: u.messagesChan}
	ctx := context.Background()
	topic := []string{"message"}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		//select {
		//case <-signals:
		<-signals
		u.consumer.Close()
		log.Fatal()
		//}
	}()

	go func() {
		for {
			for err := range u.consumer.Errors() {
				log.Error(err)
			}
		}
	}()

	go func() {
		for {
			err := u.consumer.Consume(ctx, topic, &handler)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
