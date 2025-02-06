-- name: GetAll :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name,age,gender) VALUES ($1,$2,$3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET name = $1, age = $2, gender = $3 WHERE id = $4 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
