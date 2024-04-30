-- Modify "undelegates" table
ALTER TABLE "undelegates" DROP COLUMN "user_balance", ADD COLUMN "staked_balance" character varying NULL;
