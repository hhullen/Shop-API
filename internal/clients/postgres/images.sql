-- name: AddImage :one
INSERT INTO images (uid, image)
VALUES ($1, $2)
ON CONFLICT (uid) DO NOTHING
RETURNING uid;

-- name: UpdateImage :one
UPDATE images
SET image = $1
WHERE uid = $2
RETURNING uid;

-- name: DeleteImage :one
DELETE FROM images
WHERE uid = $1
RETURNING uid;

-- name: GetProductImage :one
SELECT * FROM images i
WHERE i.uid = (SELECT image_id FROM products p WHERE p.uid = $1);

-- name: GetImage :one
SELECT image FROM images
WHERE uid = $1;