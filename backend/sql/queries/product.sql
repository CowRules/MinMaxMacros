-- name: GetProducts :many
SELECT id, name, calories, protein, fiber, price, weight, categories, shops, user_id
FROM product
GROUP BY product.id
ORDER BY product.name;

-- name: GetProduct :one
SELECT id, name, calories, protein, fiber, price, weight, categories, shops, user_id
FROM product
WHERE id=$1;

-- name: CreateProduct :one
INSERT INTO product
(id, created_at, updated_at, name, calories, protein, fiber, price, weight, categories, shops, user_id)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id, name, calories, protein, fiber, price, weight, categories, shops, user_id;

-- name: UpdateProduct :one
UPDATE product
SET updated_at=NOW(), name=$2, calories=$3, protein=$4, fiber=$5, price=$6, weight=$7, categories=$8, shops=$9
WHERE id=$1
RETURNING id, name, calories, protein, fiber, price, weight, categories, shops, user_id;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id=$1;

-- name: GetProductCategories :many
SELECT DISTINCT CAST(unnest(categories) AS TEXT) AS category
FROM product
ORDER BY category;

-- name: GetProductShops :many
SELECT DISTINCT CAST(unnest(shops) AS TEXT) AS shop
FROM product
ORDER BY shop;
