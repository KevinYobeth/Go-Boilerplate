package main

import (
	"go-boilerplate/config"
	"go-boilerplate/ports/grpc"
	"go-boilerplate/ports/http"
	"go-boilerplate/ports/scheduler"
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
