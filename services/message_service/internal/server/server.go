package server

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"
)

type messageServiceGRPCServer struct {
	messageService.UnimplementedMessageServiceServer
}

// GetRealTimeDirectMessages returns all messages in last 1 hour
func (s *messageServiceGRPCServer) GetRealTimeDirectMessages(ctx context.Context, request *messageService.GetRealTimeDirectMessagesRequest) (*messageService.GetRealTimeDirectMessagesResponse, error) {
	log.Println("GetRealTimeDirectMessages called")
	directMessage := &messageService.DirectMessage{
		id:            &wrappers.StringValue{Value: id},
		senderId:      &wrappers.StringValue{Value: senderId},
		text:          &wrappers.StringValue{Value: text},
		eventType:     &wrappers.StringValue{Value: eventType},
		createdAt:     &wrappers.StringValue{Value: createdAt},
		messageSource: &messageService.MessageSource.TWITTER,
	}
	response := []*messageService.DirectMessage{directMessage}
	return &messageService.GetRealTimeDirectMessagesResponse{response}, nil
}

// GetHistoricalDirectMessages returns all the historical messages with time frame and limit
func (s *messageServiceGRPCServer) GetHistoricalDirectMessages(ctx context.Context, request *messageService.GetHistoricalDirectMessagesRequest) (*messageService.GetHistoricalDirectMessagesResponse, error) {
	log.Println("GetHistoricalDirectMessages called")
	directMessage := &messageService.DirectMessage{
		id:            &wrappers.StringValue{Value: id},
		senderId:      &wrappers.StringValue{Value: senderId},
		text:          &wrappers.StringValue{Value: text},
		eventType:     &wrappers.StringValue{Value: eventType},
		createdAt:     &wrappers.StringValue{Value: createdAt},
		messageSource: &messageService.MessageSource.TWITTER,
	}
	response := []*messageService.DirectMessage{directMessage}
	return &messageService.GetHistoricalDirectMessagesResponse{response}, nil
}
