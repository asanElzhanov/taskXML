package services

import (
	"encoding/xml"
	"fmt"
	"server2/internal/entities"
	"server2/internal/repositories"
)

type XMLProcessor struct {
	repo *repositories.FileRepository
}

func NewXMLProcessor(fileRepository *repositories.FileRepository) *XMLProcessor {
	return &XMLProcessor{repo: fileRepository}
}

func (processor *XMLProcessor) GenerateResponse(req *entities.RequestXML) ([]byte, error) {
	response, err := processor.repo.ReadXMLFile()
	if err != nil {
		return nil, err
	}
	response.Body.Info = fmt.Sprintf("your login: %s, and password: %s.", req.Header.Login, req.Header.Password)
	updatedXML, err := xml.MarshalIndent(response, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response XML: %v", err)
	}

	return updatedXML, nil
}
