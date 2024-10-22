package database

const (
	PostOneAccountQuery = `
		INSERT INTO users (email, password, first_name, last_name, phone_number, created_by, created_time)
		VALUES ($1, $2, $3, $4, $5, $6, NOW()) 
		RETURNING id
	`

	GetAllUsers = `
		SELECT id, photo, first_name, last_name, username, email, gender, address, phone_number, 
		       email_verified_at, remember_token, created_by, updated_by, created_time, updated_time, status
		FROM users
		WHERE status = 1 
	`
)
