-- name: GetUser :one
SELECT
    username,
    password
FROM
    users
WHERE
    id = ?
    and is_deleted = false;