-- Create "undelegates" table
CREATE TABLE "undelegates" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "delegator" character varying NOT NULL, "validator" character varying NULL, "name" character varying NULL, "user_balance" character varying NULL, "amount" character varying NULL, "created_at" timestamptz NULL, "status" character varying NULL, "deploy" character varying NULL, "user_undelegates" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "undelegates_users_undelegates" FOREIGN KEY ("user_undelegates") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);