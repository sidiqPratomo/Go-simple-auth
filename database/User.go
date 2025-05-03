package database

const (
	UpdateUserQuery = `UPDATE users SET 
			status_otp = ?,
			nik = ?,
			photo = ?,
			first_name = ?, 
			last_name = ?, 
			username = ?, 
			email = ?, 
			gender = ?, 
			address = ?, 
			phone_number = ?, 
			status = ?, 
			updated_by = ?,
			updated_time = NOW()
		WHERE id = ?`

	FindUserByIdQuery = `
		SELECT 
			u.id, 
			u.status_otp,
			u.nik,
       		COALESCE(u.photo, '') AS photo,
       		u.first_name,
			u.last_name,
			u.username,
			u.email,
			COALESCE(u.gender,'') AS gender,
			COALESCE(u.address, '') AS address,
			COALESCE(u.phone_number, '') AS phone_number,
			CAST(u.email_verified_at AS DATETIME) AS email_verified_at,
			u.status,
			u.created_by,
			u.updated_by,
			u.created_time,
			u.updated_time
		FROM users u
		WHERE u.id LIKE ? 
		AND u.status = 1;
	`

	FindPrivilegeQuery = `
		SELECT rp.id, rp.role, rp.action, rp.uri, rp.method, rp.created_by ,rp.updated_by , rp.created_time , rp.updated_time, rp.status 
		FROM role_privileges rp
		WHERE rp.role IN (SELECT roles_id FROM role_users WHERE users_id = ?);
	`

	FindRoleQuery = `
		SELECT 
		ru.id, ru.users_id, 
		r.id, r.name, r.code, r.created_by, r.updated_by, r.created_time, r.updated_time, r.status, 
		ru.created_by, ru.updated_by, ru.created_time, ru.updated_time, ru.status
		FROM role_users ru
		JOIN roles r ON ru.roles_id = r.id
		WHERE ru.users_id = ?
	`

	InserOtpQuery = `
		INSERT INTO users_otps (user_id, otp, expired_at) 
		VALUES (?, ?, ?)
	`

	FindUserOtp =`
		SELECT otp FROM users_otps WHERE user_id = ? AND otp = ? AND expired_at > NOW()
	`

	FindUserOtpsByOTP =`
		SELECT id, user_id, otp, expired_at from users_otps where otp = ? and user_id = ? and expired_at > now() 
	`

	FindAccountByUserIdQuery = `
		SELECT 
			u.id, 
			u.nik,
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
			r.code as role_code,
			u.status
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.id LIKE ? 
		AND u.status = 1;
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
			r.code as role_code,
			u.status
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.email LIKE ?
		AND u.status = 1;
	`

	FindAccountByUsernameQuery = `
		SELECT 
			u.id, 
			u.nik,
			u.status_otp,
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
			r.code as role_code,
			u.status
		FROM users u
		JOIN role_users ru ON u.id = ru.users_id
		JOIN roles r ON ru.roles_id = r.id
		WHERE u.username = ? 
		AND u.status = 1
	`

	PostOneAccountQuery = `
    	INSERT INTO users (nik, email, gender, username, password, first_name, last_name, phone_number, created_time, status_otp)
    	VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), 0)
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
