-- name: CreateShippingAddress :one
INSERT INTO shipping_address (
    name, address
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetShippingAddress :one
SELECT * FROM shipping_address
WHERE id = ? LIMIT 1;

-- name: ListShippingAddress :many
SELECT * FROM shipping_address;

-- name: UpdateShippingAddress :one
UPDATE shipping_address
SET name = ?,
address = ?
WHERE id = ?
RETURNING *;

-- name: DeleteShippingAddress :exec
DELETE FROM shipping_address
WHERE id = ?;
