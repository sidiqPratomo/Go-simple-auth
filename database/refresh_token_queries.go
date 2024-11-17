package database

const (
	PostOneRefreshTokenQuery = `
		INSERT 
    	INTO oauth_refresh_tokens (id, account_id, access_token_id, revoked, expires_at)
    	VALUES (?, ?, ?, 1, NOW() + INTERVAL 1 DAY)
	`

	InvalidateRefreshTokensQuery = `
		UPDATE oauth_refresh_tokens
		SET expires_at = NOW()
		WHERE account_id = ?
	`

	FindOneRefreshTokenQuery = `
		SELECT access_token_id
		FROM oauth_refresh_tokens
		WHERE access_token_id = ?
		AND expires_at > NOW()
	`
)
