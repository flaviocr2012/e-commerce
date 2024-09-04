package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
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
		var addressJSON []byte
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &addressJSON, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Log do valor armazenado no campo address
		log.Printf("addressJSON: %s", string(addressJSON))

		// Tente decodificar como objeto único
		if err := json.Unmarshal(addressJSON, &user.Address); err != nil {
			// Log se houver falha na decodificação
			log.Printf("Falha ao decodificar address como objeto: %v", err)
			// Tente decodificar como array de endereços
			var addresses []model.Address
			if err := json.Unmarshal(addressJSON, &addresses); err != nil {
				return nil, err
			}
			if len(addresses) > 0 {
				user.Address = addresses[0]
			}
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository *UserRepositoryImpl) FindByID(id int64) (model.User, error) {
	var user model.User
	err := repository.DB.QueryRow(queries.FindUserByID, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Save(user model.User) (int64, error) {
	addressJSON, err := json.Marshal(user.Address)
	if err != nil {
		return 0, err
	}

	var id int64
	err = repository.DB.QueryRow(
		queries.InsertUser,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		addressJSON, // Passa o JSON do endereço
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
		user.Password,
		user.Role,
		user.Address,
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
