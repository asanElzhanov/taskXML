package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"server1/internal/entities"
	"server1/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{repo: repository}
}

func (service *UserService) ProcessRequest(req *entities.RequestXML) (*entities.ResponseXML, error) {
	if !service.repo.ValidateCredentials(req.Header.Login, req.Header.Password) {
		return nil, fmt.Errorf("invalid creditals")
	}
	address, err := service.repo.RetrieveAddress(req.Header.Login)
	if err != nil {
		return nil, fmt.Errorf("problem with retrieving address")
	}
	req.Header.Hash = generateHash(req.Body.Message)
	resp, err := http.Post(address, "application/xml", bytes.NewBuffer(prepareRequestXML(req)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return decodeResponseXML(resp.Body)
}

func generateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func prepareRequestXML(req *entities.RequestXML) []byte {
	modifiedXML, err := xml.Marshal(req)
	if err != nil {
		fmt.Println("Error while marshaling request")
		return nil
	}
	return modifiedXML
}

func decodeResponseXML(body io.Reader) (*entities.ResponseXML, error) {
	var recipientResponse entities.ResponseXML
	respBody, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("error while reading body")
	}
	err = xml.Unmarshal(respBody, &recipientResponse)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling XML")
	}
	recipientResponse.Header.Hash = generateHash(recipientResponse.Body.Message)
	return &recipientResponse, nil
}
