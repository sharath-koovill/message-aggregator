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

func PulsarConsumer(client pulsar.Client, topic string) string {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	var pulsarResponse string

	for i := 0; i < 10; i++ {
		// may block here
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		pulsarResponse = string(msg.Payload())

		log.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}

	return pulsarResponse
}
