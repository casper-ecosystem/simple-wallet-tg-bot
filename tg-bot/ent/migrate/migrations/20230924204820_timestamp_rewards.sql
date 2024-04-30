-- Modify "rewards_data" table
ALTER TABLE "rewards_data" ADD COLUMN "first_era_timestamp" character varying NOT NULL, ADD COLUMN "last_era_timestamp" character varying NOT NULL;
