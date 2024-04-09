package table

import "database/sql"

func verifyTable(db *sql.DB, tableName string) (bool, error) {
	query := "select tablename from pg_catalog.pg_tables where tablename = $1"

	var existingTableName string
	err := db.QueryRow(query, tableName).Scan(&existingTableName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil

}

func isFKExists(db *sql.DB, tableName, fk string) (bool, error) {
	query := `
		select count(*)
		from information_schema.table_constraints
		where table_name = $1
		and constraint_name = $2
		and constrait_type = 'FOREIGN KEY'`

	var count int
	err := db.QueryRow(query, tableName, fk).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil
}
