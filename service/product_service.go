package service

import (
	"e-commerce/model"
)

type ProductService interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id int64) (model.Product, error)
	CreateProduct(product model.Product) (int64, error)
	UpdateProduct(product model.Product) (int64, error)
	DeleteProduct(id int64) (int64, error)
}
