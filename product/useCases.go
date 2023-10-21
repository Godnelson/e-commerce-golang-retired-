package product

import (
	"encoding/json"
	"golang/config/db"
)

type productUseCasesModel struct {
	findAll  func() (*[]Product, error)
	findById func(id string) (*Product, error)
	create   func(body []byte) (*Product, error)
	update   func(body []byte) (*Product, error)
}

func UseCases() productUseCasesModel {
	dbInstance := db.Database{}.Init()
	return productUseCasesModel{
		findAll: func() (*[]Product, error) {
			var products []Product

			result := dbInstance.Find(&products)

			return &products, result.Error
		},
		findById: func(id string) (*Product, error) {
			var product Product
			result := dbInstance.Where("id = ?", id).First(&product)
			return &product, result.Error
		},
		create: func(body []byte) (*Product, error) {
			var product Product

			err := json.Unmarshal(body, &product)

			if err != nil {
				return nil, err
			}

			result := dbInstance.Create(&product)

			return &product, result.Error
		},
		update: func(body []byte) (*Product, error) {
			var product Product

			err := json.Unmarshal(body, &product)

			if err != nil {
				return nil, err
			}

			result := dbInstance.Updates(&product)

			return &product, result.Error
		},
	}
}
