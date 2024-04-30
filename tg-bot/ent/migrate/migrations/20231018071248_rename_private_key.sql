-- Modify "private_keys" table
ALTER TABLE "private_keys" DROP COLUMN "privat_key", ADD COLUMN "private_key" bytea NULL;
