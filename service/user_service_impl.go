package service

import (
	"e-commerce/model"
	"e-commerce/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// NewUserService returns a new instance of UserServiceImpl.
func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: repo,
	}
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	return s.userRepository.FindAll()
}

func (s *UserServiceImpl) GetUserByID(id int64) (model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserServiceImpl) CreateUser(user model.User) (int64, error) {
	return s.userRepository.Save(user)
}

func (s *UserServiceImpl) UpdateUser(user model.User) (int64, error) {
	return s.userRepository.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id int64) (int64, error) {
	return s.userRepository.Delete(id)
}
