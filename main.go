package main

import (
	"go-boilerplate/config"
	"go-boilerplate/src/ports"
)

func main() {
	config := config.InitConfig()

	switch config.Server.ServerType {
	case "http":
		ports.RunHTTPServer()
		return
	default:
		panic("Invalid server type")
	}
}
