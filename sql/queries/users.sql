-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, email, hashed_password
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: ResetUserTable :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUserById :one
SELECT
    *
FROM 
    users
WHERE
    email = $1;

-- name: SetUserToken :one
UPDATE users
SET access_token = $1
WHERE email = $2
RETURNING *;
