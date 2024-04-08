package kafka

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// Consumer represents a Kafka consumer.
type Consumer struct {
	client     sarama.ConsumerGroup
	topic      string
	handler    func([]byte) error
	consumerID string
}

// NewConsumer creates a new Kafka consumer instance.
func NewConsumer(brokers []string, groupID, topic, consumerID string, handler func([]byte) error) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0                                       // Set the Kafka version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // Set the rebalance strategy

	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("error creating Kafka consumer group: %w", err)
	}

	consumer := &Consumer{
		client:     client,
		topic:      topic,
		handler:    handler,
		consumerID: consumerID,
	}

	return consumer, nil
}

// Consume starts consuming messages from the Kafka topic.
func (c *Consumer) Consume(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.client.Consume(ctx, []string{c.topic}, c); err != nil {
				fmt.Printf("Error from consumer: %v\n", err)
				return
			}
			// Check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Println("Consumer terminated")
	case <-sigterm:
		fmt.Println("Terminating: signal received")
	}
	cancel()
	wg.Wait()
	return nil
}

// Setup is called on starting a new consumer group session.
func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("Consumer %s is setting up\n", c.consumerID)
	return nil
}

// Cleanup is called on closing a consumer group session.
func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("Consumer %s is cleaning up\n", c.consumerID)
	return nil
}

// ConsumeClaim is called when a new claim session is established.
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("Consumer %s is consuming from topic %s\n", c.consumerID, c.topic)
	for message := range claim.Messages() {
		err := c.handler(message.Value)
		if err != nil {
			fmt.Printf("Error handling message: %v\n", err)
		} else {
			fmt.Printf("Processed message: %s\n", string(message.Value))
			session.MarkMessage(message, "")
		}
	}
	return nil
}
