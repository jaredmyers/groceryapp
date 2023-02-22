package storage

import "github.com/jaredmyers/groceryapp/backend/models"

type Storage interface {
	GetGroceryPageProducts() (*[]models.GroceryPageProduct, error)
	GroceryListSelections() *models.GroceryListSelections
}
