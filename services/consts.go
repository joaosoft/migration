package services

const (
	DefaultURL = "http://localhost:8001"

	CREATE_MIGRATION_TABLES = `
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mode') THEN
        CREATE TYPE mode AS ENUM ('database', 'rabbitmq');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS migration (
  id_migration      TEXT NOT NULL,
  mode           	mode,
  "user"            TEXT DEFAULT user,
  executed_at       TIMESTAMP DEFAULT NOW(),
  CONSTRAINT migration_id_pkey PRIMARY KEY (id_migration, mode)
);
`
)
