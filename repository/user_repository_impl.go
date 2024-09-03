package repository

import "e-commerce/model"

type UserRepositoryImpl struct {
	// Database connection or ORM instance

}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) FindAll() ([]model.User, error) {
	// Your code here
	return nil, nil
}

func (repository UserRepositoryImpl) FindByID(id int64) (model.User, error) {
	// Your code here
	return model.User{}, nil
}

func (repository UserRepositoryImpl) Save(user model.User) (int64, error) {
	// Your code here
	return 0, nil
}

func (repository UserRepositoryImpl) Update(user model.User) (int64, error) {
	// Your code here
	return 0, nil
}

func (repository UserRepositoryImpl) Delete(id int64) (int64, error) {
	// Your code here
	return 0, nil
}
