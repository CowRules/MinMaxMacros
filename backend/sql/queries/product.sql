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
