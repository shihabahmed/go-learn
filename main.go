package main

import (
	"fmt"

	"github.com/shihabahmed/go-learn/internal/initializers"
	"github.com/shihabahmed/go-learn/internal/server"
)



func main() {
	options, _ := initializers.Run()

	server := server.Init()
	server.Run(fmt.Sprintf(":%d", options.Port))
}