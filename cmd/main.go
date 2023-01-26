package main

import (
	"log"
	qsoft_golang_entering_task "qsoft-golang-entering-task"
	"qsoft-golang-entering-task/pkg/handler"
	"qsoft-golang-entering-task/pkg/service"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(qsoft_golang_entering_task.Server)
	log.Print("Starting server...")

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running hhtp server: %s", err.Error())
	}
}
