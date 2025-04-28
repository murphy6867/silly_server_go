-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, email, hashed_password
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: ResetUserTable :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUserByEmail :one
SELECT
    *
FROM 
    users
WHERE
    email = $1;

-- name: UpdateEmailAndPassword :exec
UPDATE users
SET email = $2, hashed_password = $3, updated_at = $4
WHERE id = $1;

-- name: UpdateIsChirpsRedStatus :exec
UPDATE users
SET is_chirpy_red = $2
WHERE id = $1;