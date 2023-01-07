package routes

import (
	"context"
	"errors"
	"log"

	"github.com/jaredmyers/groceryapp/product_service/consumer/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	ctx               context.Context
	productCollection *mongo.Collection
}

func NewProductService(ctx context.Context, productCollection *mongo.Collection) *ProductService {
	return &ProductService{
		ctx:               ctx,
		productCollection: productCollection,
	}
}

func (s *ProductService) GetProducts() ([]*models.Products, error) {
	log.Println("product service getproducts is running")
	var products []*models.Products
	cursor, err := s.productCollection.Find(s.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(s.ctx) {
		var product models.Products
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(s.ctx)

	if len(products) == 0 {
		return nil, errors.New("No Documents Found")
	}

	return products, nil
}
