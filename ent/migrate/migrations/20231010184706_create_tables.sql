-- Create "organization" table
CREATE TABLE IF NOT EXISTS "organization" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" timestamptz NOT NULL, "date_updated" timestamptz NOT NULL, "public_id" character varying NOT NULL, "type" character varying NOT NULL DEFAULT 'organization', "name_dut" character varying NULL, "name_eng" character varying NULL, "identifier" jsonb NULL, "parent_id" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "organization_organization_children" FOREIGN KEY ("parent_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "organization_parent_id" to table: "organization"
CREATE INDEX IF NOT EXISTS "organization_parent_id" ON "organization" ("parent_id");
-- Create index "organization_public_id_key" to table: "organization"
CREATE UNIQUE INDEX IF NOT EXISTS "organization_public_id_key" ON "organization" ("public_id");
-- Create index "organization_type" to table: "organization"
CREATE INDEX IF NOT EXISTS "organization_type" ON "organization" ("type");
-- Create "person" table
CREATE TABLE IF NOT EXISTS "person" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" timestamptz NOT NULL, "date_updated" timestamptz NOT NULL, "public_id" character varying NOT NULL, "active" boolean NOT NULL DEFAULT false, "birth_date" character varying NULL, "email" character varying NULL, "identifier" jsonb NULL, "given_name" character varying NULL, "name" character varying NULL, "family_name" character varying NULL, "job_category" jsonb NULL, "preferred_given_name" character varying NULL, "preferred_family_name" character varying NULL, "honorific_prefix" character varying NULL, "role" jsonb NULL, "settings" jsonb NULL, "object_class" jsonb NULL, "expiration_date" character varying NULL, "token" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "person_active" to table: "person"
CREATE INDEX IF NOT EXISTS "person_active" ON "person" ("active");
-- Create index "person_email" to table: "person"
CREATE INDEX IF NOT EXISTS "person_email" ON "person" ("email");
-- Create index "person_family_name" to table: "person"
CREATE INDEX IF NOT EXISTS "person_family_name" ON "person" ("family_name");
-- Create index "person_given_name" to table: "person"
CREATE INDEX IF NOT EXISTS "person_given_name" ON "person" ("given_name");
-- Create index "person_name" to table: "person"
CREATE INDEX IF NOT EXISTS "person_name" ON "person" ("name");
-- Create index "person_public_id_key" to table: "person"
CREATE UNIQUE INDEX IF NOT EXISTS "person_public_id_key" ON "person" ("public_id");
-- Create "organization_person" table
CREATE TABLE IF NOT EXISTS "organization_person" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" timestamptz NOT NULL, "date_updated" timestamptz NOT NULL, "from" timestamptz NOT NULL, "until" timestamptz NOT NULL, "person_id" bigint NOT NULL, "organization_id" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "organization_person_organization_organizations" FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "organization_person_person_people" FOREIGN KEY ("person_id") REFERENCES "person" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "organizationperson_organization_id" to table: "organization_person"
CREATE INDEX IF NOT EXISTS "organizationperson_organization_id" ON "organization_person" ("organization_id");
-- Create index "organizationperson_person_id" to table: "organization_person"
CREATE INDEX IF NOT EXISTS "organizationperson_person_id" ON "organization_person" ("person_id");
-- Create index "organizationperson_person_id_organization_id" to table: "organization_person"
CREATE UNIQUE INDEX IF NOT EXISTS "organizationperson_person_id_organization_id" ON "organization_person" ("person_id", "organization_id");