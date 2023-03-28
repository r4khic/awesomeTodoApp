package main

import (
	todo "awesomeTodoApp/pkg"
	"awesomeTodoApp/pkg/handler"
	"awesomeTodoApp/pkg/repository"
	"awesomeTodoApp/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("[main]: Error occrured while runnig http server: %s", err.Error())
	}
}
