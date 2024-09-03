package queries

const (
	FindAllUsers = `
        SELECT id, name, email FROM users;
    `

	FindUserByID = `

        SELECT id, name, email FROM users WHERE id = $1;
    `

	InsertUser = `
        INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;
    `

	UpdateUser = `
        UPDATE users SET name = $1, email = $2 WHERE id = $3;
    `

	DeleteUser = `
        DELETE FROM users WHERE id = $1;
    `
)
