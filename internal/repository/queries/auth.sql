-- name: CreateOtp :execresult
INSERT into
    otp (phone_number, otp, expired_date)
VALUES
    (?, ?, ?);

-- name: UpdateOtp :execresult
UPDATE
    otp
SET
    otp = ?,
    expired_date = ?
WHERE
    phone_number = ?;

-- name: GetPhoneNumber :one
SELECT
    id, phone_number, otp, expired_date
FROM
    otp
WHERE
    phone_number = ?;