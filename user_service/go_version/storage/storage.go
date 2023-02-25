package storage

import (
	"database/sql"

	"github.com/jaredmyers/groceryapp/user_service/go_version/models"
)

type Storage interface {
	AddNewSelections(*models.Selections) error
}

type MySqlStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySqlStore, error) {
	// mysql setup here
	return nil, nil
}

func (m *MySqlStore) AddNewSelections(selections *models.Selections) error {
	// add selections to db here
	return nil
}
