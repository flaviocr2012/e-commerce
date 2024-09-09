package repository

import (
	"e-commerce/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int64) (model.Product, error)
	Save(product model.Product) (int64, error)
	Update(product model.Product) (int64, error)
	Delete(id int64) (int64, error)
}
