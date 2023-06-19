package pulsar

import (
	"context"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestGetPulsarClient(t *testing.T) {
	// Test that a pulsar client is returned
	client := GetPulsarClient("pulsar://localhost:6650")
	if client == nil {
		t.Error("Expected a pulsar client, got nil")
	}
}

func TestPulsarConsumer(t *testing.T) {
	// Test that a message is received from the pulsar topic
	client := GetPulsarClient("pulsar://localhost:6650")
	topic := "my-topic"
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		t.Fatalf("Could not create producer: %v", err)
	}

	message := "Hello, Pulsar!"
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(message),
	})
	if err != nil {
		t.Fatalf("Could not send message: %v", err)
	}

	response := PulsarConsumer(client, topic)
	if response != message {
		t.Errorf("Expected response '%s', got '%s'", message, response)
	}
}
