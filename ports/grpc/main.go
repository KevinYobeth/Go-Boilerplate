package grpc

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/log"
	authorsTransport "go-boilerplate/src/authors/infrastructure/transport"
	authorsService "go-boilerplate/src/authors/services"
	"net"

	GoogleGRPC "google.golang.org/grpc"
)

func RunGRPCServer() {
	logger := log.InitLogger()
	config := config.LoadServerConfig()

	logger.Info("Starting GRPC server on port " + config.ServerHost + ":" + config.ServerGRPCPort)

	srv, err := net.Listen("tcp", config.ServerHost+":"+config.ServerGRPCPort)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	server := GoogleGRPC.NewServer()

	authorsService := authorsService.NewAuthorService()
	authorsServer := authorsTransport.NewAuthorsGRPCServer(&authorsService)

	authorsServer.RegisterGRPCRoutes(server)

	logger.Fatal(server.Serve(srv))
}
