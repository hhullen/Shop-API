-- name: InsertAddress :one
INSERT INTO addresses (country, city, street)
VALUES ($1, $2, $3) ON CONFLICT (country, city, street)
DO UPDATE SET 
    country = EXCLUDED.country
RETURNING id;

-- name: InsertClient :one
INSERT INTO clients (uid, client_name, client_surname, birthday, gender, registration_date, address_id)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
ON CONFLICT (uid)
DO NOTHING
RETURNING uid;

-- name: DeleteClient :one
DELETE FROM clients WHERE uid = $1
RETURNING address_id;

-- name: CalculateClientsWithAddress :one
SELECT COUNT(*) as clientsAmount 
FROM clients
WHERE address_id = $1;

-- name: DeleteAddress :exec
DELETE FROM addresses
WHERE id = $1;

-- name: GetAllClients :many
SELECT *
FROM client_details
ORDER BY uid;

-- name: GetClientsPage :many
SELECT *
FROM client_details
ORDER BY uid
OFFSET $1
LIMIT $2;

-- name: GetClientsWithName :many
SELECT *
FROM client_details
WHERE client_name = $1 AND client_surname = $2;

-- name: UpdateClientAddress :one
UPDATE addresses
SET country = $1, city = $2, street = $3
WHERE (SELECT address_id FROM clients WHERE uid = $4) = id
RETURNING id;
