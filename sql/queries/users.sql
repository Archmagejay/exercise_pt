-- name: ListUsers :many
-- Select all rows from 'users'
SELECT * FROM users;

-- name: GetUserById :one
-- Select the specified user by UUID from 'users'
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByName :one
-- Select the specified user by name for 'users'
SELECT * FROM users
WHERE name = $1;

-- name: NewUser :one
-- Add a new user entry in 'users'
INSERT INTO users (id, name, height, start_date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteAllUsers :exec
-- Remove all users in 'users'
DELETE FROM users;

-- name: DeleteUser :exec
-- Remove the specified user from 'users'
DELETE FROM users
WHERE name = $1;