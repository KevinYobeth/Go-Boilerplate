package main

import (
	"fmt"
	"go-boilerplate/config"
)

func main() {
	config := config.InitConfig()

	fmt.Println("CONFIG", config)
}
