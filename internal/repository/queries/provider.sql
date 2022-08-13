-- name: GetProvider :many
SELECT
    id,
    name,
    is_selected,
    is_selected
FROM
    provider
WHERE
    is_deleted = false;

-- name: UpdateProvider :exec
UPDATE
    provider
SET
    is_selected = true
WHERE
    id = ?
and is_deleted = false;

-- name: UpdateFalseProvider :exec
UPDATE
    provider
SET
    is_selected = false
WHERE
    is_selected = true
and is_deleted = false;