-- name: CreateUser :exec
INSERT INTO "users" (id, name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
