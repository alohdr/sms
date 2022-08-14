-- name: CreatSms :execresult
INSERT INTO sms_history (
    id,
    sender,
    phone_number ,
    type,
    status
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: GetSmsHistory :many
SELECT
    id,
    sender,
    phone_number ,
    type,
    status
FROM sms_history
WHERE is_deleted = false
LIMIT 10
;

-- name: GetIsSelectedProvider :one
SELECT
    id,
    name
FROM provider
WHERE is_deleted = false
and is_selected = false
LIMIT 1
;