package main

import (
	"go-boilerplate/config"
	"go-boilerplate/ports/grpc"
	"go-boilerplate/ports/http"
)

func main() {
	config := config.InitConfig()

	switch config.Server.ServerType {
	case "http":
		http.RunHTTPServer()
		return
	case "grpc":
		grpc.RunGRPCServer()
	default:
		panic("Invalid server type")
	}
}
