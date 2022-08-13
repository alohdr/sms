-- name: GetProvider :one
SELECT
    id,
    name
FROM
    provider
WHERE
    id = "1"
    and is_deleted = false;

-- name: UpdateProvider :exec
UPDATE
    provider
SET
    name = ?
WHERE
    id = "1"
    and is_deleted = false;