package queries

const (
	FindAllUsers = `
		SELECT id, name, email, password, role, address, created_at, updated_at
		FROM users
		ORDER BY id
	`

	FindUserByID = `
		SELECT id, name, email, password, role, address, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	InsertUser = `
		INSERT INTO users (name, email, password, role, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id
	`

	UpdateUser = `
		UPDATE users
		SET name = $1,
		    email = $2,
		    role = $3,
		    address = $4,
		    updated_at = NOW()
		WHERE id = $5
	`

	DeleteUser = `
		DELETE FROM users
		WHERE id = $1
	`
)