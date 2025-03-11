CREATE TABLE IF NOT EXISTS "users"
(
    "id"         uuid PRIMARY KEY NOT NULL,
    "name"       varchar          NOT NULL,
    "email"      varchar UNIQUE   NOT NULL,
    "password"   varchar          NOT NULL,
    "created_at" timestamp DEFAULT 'now()',
    "updated_at" timestamp DEFAULT 'now()',
    "deleted_at" timestamp
);
