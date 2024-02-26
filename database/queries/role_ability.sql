-- name: UpsertRoleAbilityByName :one
INSERT INTO role_ability(role_id, ability_id)
SELECT role.id AS role_id, ability.id AS ability_id
FROM (SELECT id FROM role r WHERE r."name" = sqlc.arg('role')) AS role
         CROSS JOIN (SELECT id FROM ability a WHERE a."name" = sqlc.arg('ability')) AS ability
ON CONFLICT(role_id, ability_id) DO UPDATE SET role_id    = EXCLUDED.role_id,
                                               ability_id = EXCLUDED.ability_id
RETURNING *;
