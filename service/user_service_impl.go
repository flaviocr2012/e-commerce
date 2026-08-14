package service

import (
	"e-commerce/model"
	"e-commerce/repository"
	"errors"
	"strings"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: repo,
	}
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	return s.userRepository.FindAll()
}

func (s *UserServiceImpl) GetUserByID(id int64) (model.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user model.User) (int64, error) {
	// Validate required fields
	if user.Name == "" {
		return 0, errors.New("name is required")
	}
	if user.Email == "" {
		return 0, errors.New("email is required")
	}
	if user.Password == "" {
		return 0, errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return 0, errors.New("password must be at least 6 characters")
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Trim spaces
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	return s.userRepository.Save(user)
}

func (s *UserServiceImpl) UpdateUser(user model.User) (int64, error) {
	if user.ID == 0 {
		return 0, errors.New("user ID is required")
	}
	return s.userRepository.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id int64) (int64, error) {
	if id == 0 {
		return 0, errors.New("invalid user ID")
	}
	return s.userRepository.Delete(id)
}