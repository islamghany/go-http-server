-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING *;
-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;