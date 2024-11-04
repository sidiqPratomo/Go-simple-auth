package database

const (
	InserOtpQuery = `
		INSERT INTO users_otps (user_id, otp, expired_at) 
		VALUES (?, ?, ?)
	`

	FindUserOtp =`
		SELECT otp FROM users_otps WHERE user_id = ? AND otp = ? AND expired_at > NOW()
	`

	FindAccountByEmailQuery = `
		SELECT 
			u.id, 
       		COALESCE(u.photo, '') AS photo,
       		u.first_name,
			u.last_name,
			u.username,
			u.email,
			COALESCE(u.gender,'') AS gender,
			COALESCE(u.address, '') AS address,
			COALESCE(u.phone_number, '') AS phone_number,
			u.password,
			CAST(u.email_verified_at AS DATETIME) AS email_verified_at,
       		ru.roles_id, 
       		r.name AS role_name, 
			r.code as role_code
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.email LIKE ?
		AND u.status = 1;
	`

	FindAccountByUsernameQuery = `
		SELECT 
			u.id, 
       		COALESCE(u.photo, '') AS photo,
       		u.first_name,
			u.last_name,
			u.username,
			u.email,
			COALESCE(u.gender,'') AS gender,
			COALESCE(u.address, '') AS address,
			COALESCE(u.phone_number, '') AS phone_number,
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

	PostRoleUserQuery = `
    	INSERT INTO role_users (users_id, roles_id, created_time, updated_time)
    	VALUES (?, ?, NOW(), NOW())
	`

	GetAllUsers = `
		SELECT id, photo, first_name, last_name, username, email, gender, address, phone_number, 
		       email_verified_at, remember_token, created_by, updated_by, created_time, updated_time, status
		FROM users
		WHERE status = 1 
	`
)
