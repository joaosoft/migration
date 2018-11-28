-- MIGRATION
CREATE TABLE IF NOT EXISTS migration (
  id_migration      TEXT NOT NULL,
  "user"            TEXT DEFAULT user,
  executed_at       TIMESTAMP DEFAULT NOW(),
  CONSTRAINT migration_id__pkey PRIMARY KEY (id_migration)
);


-- HISTORY
CREATE TABLE IF NOT EXISTS migration_history (LIKE migration);
ALTER TABLE migration_history ADD COLUMN IF NOT EXISTS _operation TEXT NOT NULL;
ALTER TABLE migration_history ADD COLUMN IF NOT EXISTS "_user" TEXT NOT NULL;
ALTER TABLE migration_history ADD COLUMN IF NOT EXISTS "_operation_at" TIMESTAMP DEFAULT NOW();

CREATE OR REPLACE FUNCTION function_migration_history() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO migration_history VALUES(OLD.*, 'D', user, now());
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO migration_history VALUES(NEW.*, 'U', user, now());
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO migration_history VALUES(NEW.*, 'I', user, now());
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trigger_migration_history
AFTER INSERT OR UPDATE OR DELETE ON migration
    FOR EACH ROW EXECUTE PROCEDURE function_migration_history();