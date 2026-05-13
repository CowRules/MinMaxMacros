-- name: CreateUser :exec
INSERT INTO users
(id, created_at, updated_at, username, email, password)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3);

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email=$1;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id=$1;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_token
(token, created_at, updated_at, user_id, expires_at)
VALUES ($1, NOW(), NOW(), $2, $3);

-- name: GetRefreshToken :one
SELECT *
FROM refresh_token
WHERE token=$1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_token
SET revoket_at=NOW(), updated_at=NOW()
WHERE token=$1;