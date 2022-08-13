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