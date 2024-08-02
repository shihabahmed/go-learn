package main

import "github.com/shihabahmed/go-learn/internal/server"

func main() {
	server := server.Init()
	server.Run()
}