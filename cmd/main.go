package main

import (
	"log"

	"github.com/12Mighty/todo-app"
	"github.com/12Mighty/todo-app/pkg/handler"
	"github.com/12Mighty/todo-app/pkg/repository"
	"github.com/12Mighty/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
