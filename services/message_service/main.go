package main

import (
	"flag"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "path/to/generated/package" // Importing the generated gRPC code

	"message_service/internal/utils"
)

// init loads environment variables from dev.env or prod.env depending on the value of TARGET_ENV
func init() {
	targetEnv := utils.LookupEnvOrString("TARGET_ENV", "DEVELOPMENT")

	if targetEnv == "DEVELOPMENT" {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Fatal("Error loading dev.env file")
		}

	} else {
		err := godotenv.Load("prod.env")
		if err != nil {
			log.Fatal("Error loading prod.env file")
		}
	}

}

func main() {
	var (
		tcpPort = flag.String("tcp_port", utils.LookupEnvOrString("TCP_PORT", ""), "TCP port")
	)
	flag.Parse()

	// Creating a gRPC server instance
	grpcServer := grpc.NewServer()
	myMessageServer := &messageServiceGRPCServer{}
	// Registering the gRPC server implementation
	pb.RegisterMessageServiceServer(grpcServer, &myMessageServer{})

	// Listening on a TCP port
	listener, err := net.Listen("tcp", ":"+*tcpPort)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("server listening at %v", listener.Addr())
	// Starting the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
