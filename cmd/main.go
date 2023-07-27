package main

import (
	"log"

	todo "github.com/eclair-desertire/todolist-go"
	"github.com/eclair-desertire/todolist-go/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting server %s", err.Error())
	}

}
