package main

import (
	"net/http"
	"server1/internal/controllers"
	"server1/internal/repositories"
	"server1/internal/services"
)

func main() {
	dbRepo := repositories.NewDatabaseRepository("task")
	userService := services.NewUserService(dbRepo)
	handler := controllers.NewRequestHandler(userService)
	http.HandleFunc("/", handler.HandleXMLRequest)
	http.ListenAndServe(":8080", nil)
}
