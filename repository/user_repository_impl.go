package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
	// Database connection or ORM instance
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) FindAll() ([]model.User, error) {
	rows, err := repository.DB.Query(queries.FindAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UserRepositoryImpl) FindByID(id int64) (model.User, error) {
	var user model.User
	err := repository.DB.QueryRow(queries.FindUserByID, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Save(user model.User) (int64, error) {
	var id int64
	err := repository.DB.QueryRow(
		queries.InsertUser,
		user.Name,
		user.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *UserRepositoryImpl) Update(user model.User) (int64, error) {
	result, err := repository.DB.Exec(
		queries.UpdateUser,
		user.Name,
		user.Email,
		user.ID,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repository *UserRepositoryImpl) Delete(id int64) (int64, error) {
	result, err := repository.DB.Exec(
		queries.DeleteUser,
		id,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
