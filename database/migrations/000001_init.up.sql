CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "ability_enum" AS ENUM ('public', 'private');

CREATE TYPE "role_enum" AS ENUM ('user', 'admin');

CREATE TABLE "credential"
(
    "id"       UUID     NOT NULL DEFAULT uuid_generate_v4(),
    "email"    VARCHAR(254) NOT NULL,
    "password" VARCHAR(72)  NOT NULL,
    "role_id"  INTEGER      NOT NULL,

    CONSTRAINT "credential_pk" PRIMARY KEY ("id")
);

CREATE TABLE "role"
(
    "id"   SERIAL      NOT NULL,
    "name" "role_enum" NOT NULL,
    "desc" TEXT,

    CONSTRAINT "role_pk" PRIMARY KEY ("id")
);

CREATE TABLE "ability"
(
    "id"   SERIAL         NOT NULL,
    "name" "ability_enum" NOT NULL,
    "desc" TEXT,

    CONSTRAINT "ability_pk" PRIMARY KEY ("id")
);

CREATE TABLE "role_ability"
(
    "role_id"    INTEGER NOT NULL,
    "ability_id" INTEGER NOT NULL,

    CONSTRAINT "roles_on_abilities_pk" PRIMARY KEY ("role_id", "ability_id")
);

CREATE UNIQUE INDEX "credential_email_key" ON "credential" ("email");

CREATE UNIQUE INDEX "role_name_key" ON "role" ("name");

CREATE UNIQUE INDEX "ability_name_key" ON "ability" ("name");

ALTER TABLE "credential"
    ADD CONSTRAINT "credential_role_id_fk" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE "role_ability"
    ADD CONSTRAINT "roles_on_abilities_role_id_fk" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE "role_ability"
    ADD CONSTRAINT "roles_on_abilities_ability_id_fk" FOREIGN KEY ("ability_id") REFERENCES "ability" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
