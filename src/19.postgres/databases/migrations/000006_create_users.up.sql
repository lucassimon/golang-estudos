
CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT public.uuid_generate_v4(),
  "name" varchar(255) NOT NULL,
  "age" SMALLINT,
  "active" BOOLEAN DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);
