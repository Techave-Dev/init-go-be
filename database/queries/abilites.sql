-- name: UpsertAbility :one
INSERT INTO ability("name", "desc")
VALUES ($1, $2)
ON CONFLICT ("name") DO UPDATE
    SET "desc" = EXCLUDED."desc"
RETURNING *;