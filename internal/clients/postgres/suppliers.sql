-- name: InsertSupplier :one
INSERT INTO suppliers (uid, name, phone_number, address_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT (name, address_id)
DO UPDATE SET
    name = EXCLUDED.name
RETURNING uid;

-- name: UpdateSupplierAddress :one
UPDATE suppliers
SET address_id = $1
WHERE uid = $2
RETURNING uid;

-- name: DeleteSupplier :one
DELETE FROM suppliers s
WHERE s.uid = $1
RETURNING s.address_id;

-- name: CalculateSuppliersWithAddress :one
SELECT COUNT(*) as suppliersAmount
FROM suppliers
WHERE address_id = $1;

-- name: GetAllSuppliers :many
SELECT *
FROM supplier_details
ORDER BY uid;

-- name: GetSuppliersPage :many
SELECT *
FROM supplier_details
ORDER BY uid
OFFSET $1
LIMIT $2;

-- name: GetSupplier :one
SELECT *
FROM supplier_details
WHERE uid = $1;
