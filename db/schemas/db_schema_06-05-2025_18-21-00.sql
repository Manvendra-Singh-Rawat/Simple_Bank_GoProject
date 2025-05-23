CREATE TABLE IF NOT EXISTS "accounts" (
 "id" bigserial PRIMARY KEY,
 "owner" varchar NOT NULL,
 "balance" bigint NOT NULL,
 "currency" varchar NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "entries" (
 "id" bigserial PRIMARY KEY,
 "account_id" bigint NOT NULL,
 "amount" bigint NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "transfers" (
 "id" bigserial PRIMARY KEY,
 "from_account_id" bigint NOT NULL,
 "to_account_id" bigint NOT NULL,
 "amount" bigint NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'most be positive';

CREATE OR REPLACE FUNCTION TransferMoney_SP(
    FromAccountID INT,
    ToAccountID INT,
    AmountToTransfer INT
)
RETURNS VOID AS $$
BEGIN
    -- Deduct from sender
    UPDATE accounts
    SET balance = balance - AmountToTransfer
    WHERE id = FromAccountID;
    
    -- Add to reciever
    UPDATE accounts
    SET balance = balance + AmountToTransfer
    WHERE id = ToAccountID;
    
    -- Insert into transfers table about the transaction
    INSERT INTO transfers (from_account_id, to_account_id, amount)
    VALUES (FromAccountID, ToAccountID, AmountToTransfer);
    
    -- Sender's entry in entries table
    INSERT INTO entries (account_id, amount)
    VALUES (FromAccountID, -AmountToTransfer);
    
    -- Reciever's entry in entries table
    INSERT INTO entries (account_id, amount)
    VALUES (ToAccountID, AmountToTransfer);
END;
$$ LANGUAGE plpgsql;
