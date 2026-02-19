-- name: InsertProduct :one
INSERT INTO products (uid, name, category, price,
    available_stock, last_update_date, supplier_id, image_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (name, category, supplier_id)
DO UPDATE SET
    price = EXCLUDED.price,
    available_stock = products.available_stock + EXCLUDED.available_stock,
    last_update_date = EXCLUDED.last_update_date
RETURNING uid;

-- name: LockStockForUpdate :one
SELECT available_stock
FROM products
WHERE uid = $1
FOR UPDATE;

-- name: DecreaseProduct :one
UPDATE products
SET available_stock = available_stock - sqlc.arg(amount)
WHERE uid = sqlc.arg(uid)
RETURNING available_stock;

-- name: GetProduct :one
SELECT *
FROM products p
WHERE p.uid = $1;

-- name: GetAllProducts :many
SELECT *
FROM products p
ORDER BY p.uid;

-- name: GetProductsPage :many
SELECT *
FROM products p
ORDER BY p.uid
OFFSET $1
LIMIT $2;

-- name: DeleteProduct :one
DELETE FROM products p
WHERE p.uid = $1
RETURNING p.uid;

-- name: IsImageAndSupplierExists :one
SELECT (
    EXISTS(SELECT 1 FROM images i WHERE i.uid = sqlc.arg(image_uid))
    AND
    EXISTS(SELECT 1 FROM suppliers s WHERE s.uid = sqlc.arg(supplier_uid))
)::bool AS is_exists;