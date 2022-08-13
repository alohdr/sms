-- name: GetUser :one
SELECT
    username,
    password
FROM
    users
WHERE
    username = ?
    and is_deleted = false;