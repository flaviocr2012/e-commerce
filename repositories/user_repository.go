package repositories

import (
	"e-commerce/entities"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByID(id int64) (entities.User, error)
	Save(user entities.User) (int64, error)
	Update(user entities.User) (int64, error)
	Delete(id int64) (int64, error)
}
