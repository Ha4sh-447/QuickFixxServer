package kafka

import (
	"UserServiceQF/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

func SetupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		return nil, fmt.Errorf("couldn't create producer: %w", err)
	}

	return producer, nil
}

func SendKafkaMessage(producer sarama.SyncProducer, order models.KafkaMsg) error {

	orderJSON, err := json.Marshal(&order)
	log.Println("USERORDER", orderJSON)
	if err != nil {
		log.Println("Kafka producer JSON error", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "User_Order",
		Key:   sarama.StringEncoder(strconv.Itoa(order.UserId)),
		Value: sarama.StringEncoder(orderJSON),
	}

	_, _, err = producer.SendMessage(msg)
	// log.Println("KAFKA PRODUCER", partition, offset, err)

	return err
}
