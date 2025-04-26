-- name: CreateChirp :one
INSERT INTO chirps (
    id, user_id, body, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAllChirp :many
SELECT
  *
FROM
  chirps
ORDER BY
  created_at ASC;

-- name: GetChirpById :one
SELECT
  *
FROM
  chirps
WHERE
  id = $1;

-- name: DeleteChirpById :exec
DELETE FROM chirps
WHERE id = $1;
