package controllers

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"server2/internal/entities"
	"server2/internal/services"
)

type RequestHandler struct {
	processor *services.XMLProcessor
}

func NewRequestHandler(proc *services.XMLProcessor) *RequestHandler {
	return &RequestHandler{
		processor: proc,
	}
}

func (handler *RequestHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var req entities.RequestXML
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		log.Printf("Error reading request body: %v", err)
		return
	}
	err = xml.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse XML", http.StatusBadRequest)
		log.Printf("Error parsing XML: %v", err)
		return
	}
	responseXML, err := handler.processor.GenerateResponse(&req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error generating response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(responseXML)
}
