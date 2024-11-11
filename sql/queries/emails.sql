-- name: CreateEmail :one
INSERT INTO emails (id, created_at, updated_at, email)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEmailByID :one
SELECT id, created_at, updated_at, email, is_verified
FROM emails
WHERE id = $1;

-- name: GetEmail :one
SELECT id, created_at, updated_at, email, is_verified
FROM emails
WHERE email = $1;

-- name: UpdateEmailIsVerified :exec
UPDATE emails
SET is_verified = $2
WHERE id = $1;