-- name: CreateUser :one
INSERT INTO users (id, email, hashedpassword, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()

)
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;


-- name: GetRoleByID :one
SELECT role FROM users WHERE id = $1;