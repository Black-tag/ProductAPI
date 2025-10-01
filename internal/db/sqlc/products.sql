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


-- name: DeleteProductByID :exec
DELETE FROM products
WHERE id = $1;


-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1;


-- name: UpdateProduct :one
UPDATE products
SET 
    name = $2,
    price = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
