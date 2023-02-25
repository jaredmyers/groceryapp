package storage

import (
	"log"

	"github.com/jaredmyers/groceryapp/user_service/go_version/models"
)

type MockStorage struct{}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (s *MockStorage) AddNewSelections(selections *models.Selections) error {

	log.Println("new selections added")
	log.Println(selections)

	return nil
}
