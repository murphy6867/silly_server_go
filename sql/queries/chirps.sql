-- name: CreateChirp :one
INSERT INTO chirps (
    id, user_id, body, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetChirpById :one
SELECT
  *
FROM
  chirps
WHERE
  id = $1;

-- name: GetAllChirpsAsc :many
SELECT * FROM chirps
ORDER BY created_at ASC;

-- name: GetAllChirpsDesc :many
SELECT * FROM chirps
ORDER BY created_at DESC;

-- name: GetChirpsByUserIdAsc :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;

-- name: GetChirpsByUserIdDesc :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: DeleteChirpById :exec
DELETE FROM chirps
WHERE id = $1;
