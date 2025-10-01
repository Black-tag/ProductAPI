-- name: CreateRefreshToken :exec
INSERT  INTO refresh_tokens (token, user_id, created_at, updated_at, expires_at, revoked_at)
VALUES ($1, $2, $3, $4, $5, $6);


-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token = $1; 