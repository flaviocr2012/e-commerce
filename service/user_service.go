package service

import "e-commerce/model"

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int64) (model.User, error)
	CreateUser(user model.User) (int64, error)
	UpdateUser(user model.User) (int64, error)
	DeleteUser(id int64) (int64, error)
}
