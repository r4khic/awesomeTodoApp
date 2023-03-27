package main

import (
	todo "awesomeTodoApp/pkg"
	"awesomeTodoApp/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("[main]: Error occrured while runnig http server: %s", err.Error())
	}
}
