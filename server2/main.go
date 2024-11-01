package main

import (
	"fmt"
	"net/http"
	"server2/internal/controllers"
	"server2/internal/repositories"
	"server2/internal/services"
)

func main() {
	repo := repositories.NewFileRepository()
	processor := services.NewXMLProcessor(repo)
	handler := controllers.NewRequestHandler(processor)
	http.HandleFunc("/", handler.HandleRequest)
	fmt.Println("Server running on port 3003")
	http.ListenAndServe(":3003", nil)
}
