-- name: GetProducts :many
SELECT id, name, calories, protein, fiber, price, weight, categories, shops
FROM product
GROUP BY product.id
ORDER BY product.name;