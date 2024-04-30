-- Modify "rewards_data" table
ALTER TABLE "rewards_data" DROP COLUMN "first_block", DROP COLUMN "last_block", ADD COLUMN "first_era" bigint NOT NULL DEFAULT 0, ADD COLUMN "last_era" bigint NOT NULL DEFAULT 0;
