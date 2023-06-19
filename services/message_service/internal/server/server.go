package server

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	messageService "github.com/example/messageService"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sharath-koovill/message-aggregator/blob/master/services/message_service/internal/pulsar/pulsar_consumer"
	"github.com/sharath-koovill/message-aggregator/blob/master/services/message_service/internal/utils"
)

type messageServiceGRPCServer struct {
	messageService.UnimplementedMessageServiceServer
}

type Message struct {
	id            string
	senderId      string
	text          string
	eventType     string
	createdAt     string
	messageSource string
}

// GetRealTimeDirectMessages returns all messages in last 1 hour
func (s *messageServiceGRPCServer) GetRealTimeDirectMessages(ctx context.Context, request *messageService.GetRealTimeDirectMessagesRequest) (*messageService.GetRealTimeDirectMessagesResponse, error) {
	log.Println("GetRealTimeDirectMessages called")

	var message Message
	pulsarMessageString := GetMessagesFromPulsar()

	err := json.Unmarshal([]byte(pulsarMessageString), &message)
	if err != nil {
		log.Fatalf("Unable to parse message string to json")
	}

	directMessage := &messageService.DirectMessage{
		id:            &wrappers.StringValue{Value: message.id},
		senderId:      &wrappers.StringValue{Value: message.senderId},
		text:          &wrappers.StringValue{Value: message.text},
		eventType:     &wrappers.StringValue{Value: message.eventType},
		createdAt:     &wrappers.StringValue{Value: message.createdAt},
		messageSource: &messageService.MessageSource.TWITTER,
	}
	response := []*messageService.DirectMessage{directMessage}
	return &messageService.GetRealTimeDirectMessagesResponse{response}, nil
}

// GetHistoricalDirectMessages returns all the historical messages with time frame and limit
func (s *messageServiceGRPCServer) GetHistoricalDirectMessages(ctx context.Context, request *messageService.GetHistoricalDirectMessagesRequest) (*messageService.GetHistoricalDirectMessagesResponse, error) {
	log.Println("GetHistoricalDirectMessages called")

	var message Message
	pulsarMessageString := GetMessagesFromPulsar()

	err := json.Unmarshal([]byte(pulsarMessageString), &message)
	if err != nil {
		log.Fatalf("Unable to parse message string to json")
	}

	directMessage := &messageService.DirectMessage{
		id:            &wrappers.StringValue{Value: message.id},
		senderId:      &wrappers.StringValue{Value: message.senderId},
		text:          &wrappers.StringValue{Value: message.text},
		eventType:     &wrappers.StringValue{Value: message.eventType},
		createdAt:     &wrappers.StringValue{Value: message.createdAt},
		messageSource: &messageService.MessageSource.TWITTER,
	}
	response := []*messageService.DirectMessage{directMessage}
	return &messageService.GetHistoricalDirectMessagesResponse{response}, nil
}

func GetMessagesFromPulsar() string {
	var (
		pulsarUrl   = flag.String("pulsar_url", utils.LookupEnvOrString("PULSAR_URL", ""), "Pulsar Url")
		pulsarTopic = flag.String("pulsar_topic", utils.LookupEnvOrString("PULSAR_TOPIC", ""), "Pulsar Topic")
	)
	flag.Parse()

	pulsarClient := pulsar_consumer.GetPulsarClient(pulsarUrl)
	return pulsar_consumer.PulsarConsumer(pulsarClient, pulsarTopic)
}
