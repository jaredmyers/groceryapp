package storage

import "github.com/jaredmyers/groceryapp/backend/models"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) GetGroceryPageProducts() (*[]models.GroceryPageProduct, error) {
	var testProducts = []models.GroceryPageProduct{
		{ID: 1, Item: "Raisin Bagel"},
		{ID: 2, Item: "Milk"},
		{ID: 3, Item: "Waffles"},
		{ID: 4, Item: "Syrup"},
		{ID: 5, Item: "Mexican Cheese"},
		{ID: 6, Item: "Chicken Breast"},
		{ID: 7, Item: "Wrap"},
		{ID: 8, Item: "Wrap"},
		{ID: 9, Item: "Hummus"},
	}
	return &testProducts, nil
}

func (s *MemoryStorage) GroceryListSelections() *models.GroceryListSelections {
	return &models.GroceryListSelections{}

}
