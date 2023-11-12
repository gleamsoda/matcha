DROP TABLE IF EXISTS "issues";
CREATE TABLE "issues" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "title" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
