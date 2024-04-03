-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;


-- name: CreateUser :execresult
INSERT INTO users (
  name, email, age
) VALUES (
  ?, ?, ?
);

-- name: UpdateUserAges :exec
UPDATE users SET age = ?
WHERE id = ?;
