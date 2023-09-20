package postgres

import (
	"database/sql"
	"fmt"

	// Register postgres driver
	_ "github.com/lib/pq"
)

func Connect(connection string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ApplyMigrations() {

}
