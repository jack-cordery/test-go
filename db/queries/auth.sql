-- name: GetUserByEmail :one
SELECT
  *
FROM
  users
WHERE
  email = $1
LIMIT
  1;

-- name: GetUserByID :one
SELECT
  *
FROM
  users
WHERE
  id = $1
LIMIT
  1;

-- name: CreateUser :one
INSERT INTO
  users (email)
VALUES
  ($1)
RETURNING
  id;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE
  id = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE
  email = $1;
