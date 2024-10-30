-- name: CreateToken :one
INSERT INTO tokens (id, email_id, token, created_at, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;