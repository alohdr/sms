-- name: GetMakingPayment :many
SELECT bank_code, type, description FROM making_payment WHERE bank_code = ?