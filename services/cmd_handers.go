package services

import "database/sql"

func MigrationHandler(option MigrationOption, tx *sql.Tx, data string) error {
	_, err := tx.Exec(data)

	return err
}
