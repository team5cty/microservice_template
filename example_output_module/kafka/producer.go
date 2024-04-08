package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/segmentio/ksuid"
)

// ProduceMessage produces a message to the Kafka topic
func ProduceMessage(topic string, message string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return fmt.Errorf("error creating Kafka producer: %s", err)
	}
	defer producer.Close()

	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(ksuid.New().String()),
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		return fmt.Errorf("error sending message to Kafka: %s", err)
	}

	return nil
}
