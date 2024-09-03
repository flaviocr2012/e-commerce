package repository

import (
	"e-commerce/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id int64) (model.User, error)
	Save(user model.User) (int64, error)
	Update(user model.User) (int64, error)
	Delete(id int64) (int64, error)
}
