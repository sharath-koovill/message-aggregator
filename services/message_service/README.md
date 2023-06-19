# Message service is a microservice that has 2 endpoints and runs on localhost:50051
- GetRealTimeDirectMessages : this endpoint will return the directmessages in the last 1 hour.
- GetHistoricalDirectMessages : this endpoint returns historical direct messages for a specific date and with limit. 

## Prerequisites for running this service
Make sure the apache pulsar is running at localhost:6650

### Generate proto stub files
`cd api`
`protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative message_service.proto`

## Command to run
### Build the message_service docker container
`docker build -t message_service .`

## Run message_service
`docker run message_service`

## Unit test
go test

