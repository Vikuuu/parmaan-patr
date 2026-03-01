-- name: CreatePaymentDetail :one
INSERT INTO payment_detail (
    acc_holder, acc_number, ifsc, branch, bank_name,
    virtual_payment_addr, fk_company_id
) VALUES (
    ?, ?, ?, ?, ?, 
    ?, ?
)
RETURNING *;

-- name: ListPaymentDetail :many
SELECT * FROM payment_detail;

-- name: GetPaymentDetailWithID :one
SELECT * FROM payment_detail
WHERE id = ? LIMIT 1;

-- name: GetPaymentDetailWithCompanyID :one
SELECT * FROM payment_detail
WHERE fk_company_id = ? LIMIT 1;

-- name: UpdatePaymentDetail :one
UPDATE payment_detail
SET acc_holder = ?,
acc_number = ?,
ifsc = ?,
branch = ?,
bank_name = ?,
virtual_payment_addr = ?,
fk_company_id = ?
WHERE id = ?
RETURNING *;

-- name: DeletePaymentDetail :exec
DELETE FROM payment_detail
WHERE id = ?;
