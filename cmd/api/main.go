package main

import (
	"todo-list/internal/server/http"
	"todo-list/pkg/config"
)

func main() {
	config.LoadConfig()

	server := http.NewServer()
	server.Run()
}
