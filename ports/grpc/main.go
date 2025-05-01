package grpc

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinyobeth/go-boilerplate/config"
	authorsGRPC "github.com/kevinyobeth/go-boilerplate/internal/authors/presentation/grpc"
	authorsService "github.com/kevinyobeth/go-boilerplate/internal/authors/services"
	"github.com/kevinyobeth/go-boilerplate/shared/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/shared/log"

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
	authorsServer := authorsGRPC.NewAuthorsGRPCServer(&authorsService)

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
