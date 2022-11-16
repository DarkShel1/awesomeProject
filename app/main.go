package main

import (
	"awesomeProject"
	"awesomeProject/handlers"
	"log"
)

func main() {
	handler := new(handlers.Handler)

	server := new(awesomeProject.Server)
	err := server.Run("8090", handler.InitRoutes())
	if err != nil {
		log.Fatal(err.Error())
	}
}
