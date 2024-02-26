ALTER TABLE "credential"
    DROP CONSTRAINT IF EXISTS "credential_role_id_fk";

ALTER TABLE "role_ability"
    DROP CONSTRAINT IF EXISTS "roles_on_abilities_role_id_fk";

ALTER TABLE "role_ability"
    DROP CONSTRAINT IF EXISTS "roles_on_abilities_ability_id_fk";

DROP TABLE IF EXISTS "role_ability";

DROP TABLE IF EXISTS "ability";

DROP TABLE IF EXISTS "role";

DROP TABLE IF EXISTS "credential";

DROP TYPE IF EXISTS "ability_enum";

DROP TYPE IF EXISTS "role_enum";