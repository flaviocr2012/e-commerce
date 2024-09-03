package repositories

import "e-commerce/entities"

type UserRepositoryImpl struct {
	// Database connection or ORM instance

}

func (repository UserRepositoryImpl) FindAll() ([]entities.User, error) {
	// Your code here
	return nil, nil
}

func (repository UserRepositoryImpl) FindByID(id int64) (entities.User, error) {
	// Your code here
	return entities.User{}, nil
}

func (repository UserRepositoryImpl) Save(user entities.User) (int64, error) {
	// Your code here
	return 0, nil
}

func (repository UserRepositoryImpl) Update(user entities.User) (int64, error) {
	// Your code here
	return 0, nil
}

func (repository UserRepositoryImpl) Delete(id int64) (int64, error) {
	// Your code here
	return 0, nil
}
