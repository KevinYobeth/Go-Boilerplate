package main

import (
	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/ports/grpc"
	"github.com/kevinyobeth/go-boilerplate/ports/http"
	"github.com/kevinyobeth/go-boilerplate/ports/scheduler"
)

func main() {
	configs := config.InitConfig()

	switch configs.Server.ServerType {
	case "http":
		http.RunHTTPServer()
		return
	case "grpc":
		grpc.RunGRPCServer()
	case "scheduler":
		scheduler.RunScheduler()
	default:
		panic("Invalid server type")
	}
}
