-- name: CreateUser :exec
INSERT INTO "users" (id, name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: GetUserByEmail :one
SELECT *
FROM "users"
WHERE email = $1
  AND deleted_at IS NULL;

-- name: GetUserByID :one
SELECT *
FROM "users"
WHERE id = $1;

-- name: UpdateUserAvatar :exec
UPDATE "users"
SET avatar_url=$2,
    updated_at=NOW()
WHERE id = $1;