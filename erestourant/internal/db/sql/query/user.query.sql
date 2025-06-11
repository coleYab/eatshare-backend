-- name: GetUserById :one
SELECT * FROM "user" WHERE id=$1;

-- name: CreateUser :one
INSERT INTO "user"("name", "email", "password")
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email=$1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "user";

-- name: DeleteUserById :exec
DELETE FROM "user" WHERE id=$1;

-- name: UpdateUser :one
UPDATE "user"
SET
    "name" = $2,
    "email" = $3,
    "password" = $4,
    "updatedAt" = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
