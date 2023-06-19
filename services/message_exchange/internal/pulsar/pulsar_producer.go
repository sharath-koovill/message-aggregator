package pulsar

import (
	"context"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func GetPulsarClient(connUrl string) pulsar.Client {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               connUrl,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	return client
}

func ProduceMessage(client pulsar.Client, topic string, message string) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(message),
	})

	defer producer.Close()

	if err != nil {
		log.Println("Failed to publish message", err)
	}
	log.Println("Published message")
}
