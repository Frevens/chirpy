-- name: GetUserFromRefreshToken :one
SELECT users.*
FROM users
JOIN refresh_tokens
    ON refresh_tokens.user_id = users.id
WHERE refresh_tokens.token = $1
    AND refresh_tokens.expires_at > NOW()
    AND refresh_tokens.revoked_at IS NULL;