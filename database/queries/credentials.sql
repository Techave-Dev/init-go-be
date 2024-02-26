-- name: InsertCredential :one
INSERT INTO credential("email", "password", "role_id")
SELECT sqlc.arg('email') as email, sqlc.arg('password') as password, "role"."id" as role_id
FROM "role"
WHERE "role"."name" = sqlc.arg('role')
RETURNING *;

-- name: FindCredentialByEmail :one
SELECT *
FROM credential
WHERE email = $1;

-- name: FindCredentialById :one
SELECT *
FROM credential
WHERE id = $1;

-- name: FindCredentialAbilities :many
SELECT a.name::ability_enum
FROM role_ability ra
         LEFT JOIN ability a on a.id = ra.ability_id
         LEFT JOIN public.role r on r.id = ra.role_id
         LEFT JOIN public.credential c on r.id = c.role_id
WHERE r.name IS NOT NULL
  AND c.id = sqlc.arg('credentialId');