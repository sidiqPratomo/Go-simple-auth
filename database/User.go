package database

const (
	InserOtpQuery = `
		INSERT INTO users_otps (user_id, otp, expired_at) 
		VALUES (?, ?, ?)
	`
	FindAccountByEmailQuery = `
		SELECT u.id, 
       		u.photo, 
       		u.first_name,
			u.last_name,
			u.username,
			u.email,
			u.gender,
			u.address,
			u.phone_number,
			u.password,
			u.email_verified_at,
       		ru.roles_id, 
       		r.name AS role_name, 
			r.code as role_code
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.email LIKE ?
		AND u.status = 1
		AND u.email_verified_at IS NOT NULL;
	`

	FindAccountByUsernameQuery = `
		SELECT u.id, 
       		u.photo, 
       		u.first_name,
			u.last_name,
			u.username,
			u.email,
			u.gender,
			u.address,
			u.phone_number,
			u.password,
			u.email_verified_at,
       		ru.roles_id, 
       		r.name AS role_name, 
			r.code as role_code
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.username LIKE ? 
		AND u.status = 1
	`

	PostOneAccountQuery = `
    	INSERT INTO users (email, gender, username, password, first_name, last_name, phone_number, created_time)
    	VALUES (?, ?, ?, ?, ?, ?, ?, NOW())
	`

	GetAllUsers = `
		SELECT id, photo, first_name, last_name, username, email, gender, address, phone_number, 
		       email_verified_at, remember_token, created_by, updated_by, created_time, updated_time, status
		FROM users
		WHERE status = 1 
	`
)
