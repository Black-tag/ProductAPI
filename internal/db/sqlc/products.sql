-- name: CreateProductsFromRequest :one
INSERT INTO products (id, name, price, created_at, updated_at, posted_by)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    $3

)
RETURNING *;


-- name: GetAllProducts :many
SELECT * FROM products;
