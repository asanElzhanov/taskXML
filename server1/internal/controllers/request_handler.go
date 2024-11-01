package controllers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"server1/internal/entities"
	"server1/internal/services"
)

type RequestHandler struct {
	service *services.UserService
}

func NewRequestHandler(s *services.UserService) *RequestHandler {
	return &RequestHandler{
		service: s,
	}
}

func (handler *RequestHandler) HandleXMLRequest(w http.ResponseWriter, r *http.Request) {
	var req entities.RequestXML
	err := xml.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode XML", http.StatusBadRequest)
		return
	}
	resp, err := handler.service.ProcessRequest(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseXML, err := xml.Marshal(resp)
	if err != nil {
		fmt.Println("error while marshali g")
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(responseXML)
}
