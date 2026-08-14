package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {
	rows, err := r.DB.Query(queries.FindAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Address, // Agora é string - sem JSON
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepositoryImpl) FindByID(id int64) (model.User, error) {
	var user model.User
	err := r.DB.QueryRow(queries.FindUserByID, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Address, // Agora é string
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(user model.User) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		queries.InsertUser,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.Address, // Agora é string
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepositoryImpl) Update(user model.User) (int64, error) {
	result, err := r.DB.Exec(
		queries.UpdateUser,
		user.Name,
		user.Email,
		user.Role,
		user.Address, // Agora é string
		user.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *UserRepositoryImpl) Delete(id int64) (int64, error) {
	result, err := r.DB.Exec(queries.DeleteUser, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}