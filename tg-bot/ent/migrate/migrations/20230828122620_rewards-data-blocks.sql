-- Modify "rewards_data" table
ALTER TABLE "rewards_data" ADD COLUMN "first_block" bigint NOT NULL DEFAULT 0, ADD COLUMN "last_block" bigint NOT NULL DEFAULT 0;
