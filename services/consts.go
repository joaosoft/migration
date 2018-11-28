package services

const (
	DefaultURL = "http://localhost:8001"

	CREATE_MIGRATION_TABLES = `
CREATE TABLE IF NOT EXISTS migration (
  id_migration      TEXT NOT NULL,
  "user"            TEXT DEFAULT user,
  executed_at       TIMESTAMP DEFAULT NOW(),
  CONSTRAINT migration_id__pkey PRIMARY KEY (id_migration)
);
`
)
