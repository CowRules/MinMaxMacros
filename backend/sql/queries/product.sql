-- name: GetProducts :many
SELECT id, name, calories, protein, fiber, price, weight, categories, shops
FROM product
GROUP BY product.id
ORDER BY product.name;

-- name: GetProductCategories :many
SELECT DISTINCT CAST(unnest(categories) AS TEXT) AS category
FROM product
ORDER BY category;

-- name: GetProductShops :many
SELECT DISTINCT CAST(unnest(shops) AS TEXT) AS shop
FROM product
ORDER BY shop;

-- name: CreateProduct :one
INSERT INTO product
(id, created_at, updated_at, name, calories, protein, fiber, price, weight, categories, shops)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, calories, protein, fiber, price, weight, categories, shops;
