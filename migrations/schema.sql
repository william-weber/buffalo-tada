CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "todos" (
"id" TEXT PRIMARY KEY,
"text" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
