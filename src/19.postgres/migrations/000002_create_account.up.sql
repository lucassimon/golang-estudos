
-- CREATE TYPE "Currency" AS ENUM (
--   'USD',
--   'EUR'
-- );

CREATE TABLE IF NOT EXISTS "accounts" (
  "id" uuid PRIMARY KEY DEFAULT public.uuid_generate_v4(),
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
