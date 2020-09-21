CREATE TABLE IF NOT EXISTS "transfers" (
  "id" uuid PRIMARY KEY DEFAULT  public.uuid_generate_v4(),
  "from_account_id" uuid NOT NULL,
  "to_account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
