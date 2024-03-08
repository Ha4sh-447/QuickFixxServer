package kafka

import (
	"OrderServiceQF/messaging"
	"OrderServiceQF/models"
	"context"
	"fmt"
	"log"

	jsoniter "github.com/json-iterator/go"

	"github.com/IBM/sarama"
)

// consumer
const (
	ConsumerGroup   = "orders-group"
	ConsumerTopic   = "User_Order"
	KafkaServerAddy = "localhost:9092"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func InitializeConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	newConsumer, err := sarama.NewConsumer([]string{KafkaServerAddy}, config)

	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group", err)
	}

	return newConsumer, nil
}

func SetupConsumerGroup(ctx context.Context, order models.UserOrderProd) {
	consumer, err := InitializeConsumer()
	if err != nil {
		log.Println("Error initalizing consumer", err)
	}
	defer consumer.Close()

	partionConsumer, err := consumer.ConsumePartition(ConsumerTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("ERROR creating partition consumer ", err)
	}

	defer func() {
		if err := partionConsumer.Close(); err != nil {
			log.Println("ERROR closing partition consumer ", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Context done. Exiting consumer loop.")
			return
		case msg := <-partionConsumer.Messages():
			// kafka dm's me
			log.Println("Received message:", string(msg.Key), string(msg.Value))
			var order models.KafkaMsg
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Println("ERROR DECODING MESSAGE", err)
			}
			err = messaging.TwilioServeSms(&order)
			if err != nil {
				log.Println("TWILIO ERROR", err)
			}
			messaging.SendMessage("", string(order.Name), string(order.DateOrdered))
		}
	}
}
