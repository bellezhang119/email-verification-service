-- name: CreateToken :one
INSERT INTO tokens (id, email_id, token, created_at, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetToken :one
SELECT id, email_id, token, created_at, expires_at
FROM tokens
WHERE token = $1;

-- name: UpdateTokenIsUsed :exec
UPDATE tokens
SET is_used = $2
WHERE token = $1;
