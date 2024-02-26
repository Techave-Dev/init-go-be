-- name: UpsertRole :one
INSERT INTO role("name", "desc")
VALUES ($1, $2)
ON CONFLICT ("name") DO UPDATE
    SET "desc" = EXCLUDED."desc"
RETURNING *;