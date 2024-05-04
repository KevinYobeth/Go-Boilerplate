package grpc

import (
	"go-boilerplate/config"
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	authorsTransport "go-boilerplate/src/authors/infrastructure/transport"
	authorsService "go-boilerplate/src/authors/services"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Serve(srv); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	<-signals

	logger.Info("Shutting down server...")
	server.GracefulStop()

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Server Shutdown")
}
