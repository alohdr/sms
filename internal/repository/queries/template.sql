-- name: CreateTemplate :execresult
INSERT INTO template (
    id, 
    type, 
    apps_name, 
    text
) VALUES (
    ?, ?, ?, ?
);

-- name: GetListTemplate :many
SELECT
    id, 
    type, 
    apps_name, 
    text
FROM template
WHERE is_deleted = false
LIMIT 10
;

-- name: UpdateTemplate :exec
UPDATE template
SET 
    type = ?,
    apps_name = ?, 
    text = ?
WHERE
    id = ?
AND
    is_deleted = false
;

-- name: DeleteTemplate :exec
UPDATE template
SET is_deleted = true
WHERE
    id = ?
AND
    is_deleted = false
;