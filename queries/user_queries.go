package queries

const (
	FindAllUsers = `
        SELECT id, name, email, password, role, address, created_at, updated_at FROM users;
    `

	FindUserByID = `
        SELECT id, name, email, password, role, address, created_at, updated_at FROM users WHERE id = $1;
    `

	InsertUser = `
        INSERT INTO users (name, email, password, role, address) VALUES ($1, $2, $3, $4, $5) RETURNING id;
    `

	UpdateUser = `
        UPDATE users SET name = $1, email = $2, password = $3, role = $4, address = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6;
    `

	DeleteUser = `
        DELETE FROM users WHERE id = $1;
    `
)
