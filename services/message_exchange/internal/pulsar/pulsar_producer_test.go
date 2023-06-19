package pulsar

import (
	"context"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestGetPulsarClient(t *testing.T) {
	// Test that a pulsar client is returned
	client := GetPulsarClient("pulsar://localhost:6650")
	if client == nil {
		t.Errorf("GetPulsarClient returned nil")
	}

	// Test that the client is of type pulsar.Client
	if _, ok := client.(pulsar.Client); !ok {
		t.Errorf("GetPulsarClient did not return a pulsar.Client")
	}
}

func TestProduceMessage(t *testing.T) {
	// Create a pulsar client
	client := GetPulsarClient("pulsar://localhost:6650")

	// Test that a message is produced successfully
	topic := "test-topic"
	message := "test-message"
	ProduceMessage(client, topic, message)

	// Consume the message to ensure it was produced successfully
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "test-subscription",
		Type:             pulsar.Shared,
	})
	if err != nil {
		t.Errorf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg, err := consumer.Receive(ctx)
	if err != nil {
		t.Errorf("Failed to consume message: %v", err)
	}

	if string(msg.Payload()) != message {
		t.Errorf("Produced message does not match consumed message")
	}
}
