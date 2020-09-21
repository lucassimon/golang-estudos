CREATE TABLE IF NOT EXISTS "entries" (
  "id" uuid PRIMARY KEY DEFAULT  public.uuid_generate_v4(),
  "account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
