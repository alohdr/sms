-- name: GetMakingPayment :one
SELECT bank_code, atm, internet_banking, mobile_banking FROM making_payment WHERE bank_code = ?
LIMIT 1
;