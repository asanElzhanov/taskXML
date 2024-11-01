package repositories

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"server2/internal/entities"
)

type FileRepository struct {
	FilePath string
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		FilePath: filepath.Join("client", "file.xml"),
	}
}

func (repo *FileRepository) ReadXMLFile() (*entities.ResponseXML, error) {
	xmlData, err := os.ReadFile(repo.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error with reading XML file, %v", err)
	}
	var response entities.ResponseXML
	err = xml.Unmarshal(xmlData, &response)
	if err != nil {
		return nil, fmt.Errorf("error with unmarshaling XML file, %v", err)
	}
	return &response, nil
}
