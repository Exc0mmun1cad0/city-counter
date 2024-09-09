package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(user, dbname, password, sslmode string) (*PostgresStorage, error) {
	// constructing connection string
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=%s",
		user, dbname, password, sslmode,
	)

	// connecting to Postgres
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres: %w", err)
	}

	// testing connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (pg *PostgresStorage) CountCitiesInCountry(countryCode string) (int, error) {
	query := `SELECT COUNT(*) from city WHERE countrycode = '%s'`

	// executing query to db
	row, err := pg.db.Query(query, countryCode)
	if err != nil {
		return 0, fmt.Errorf("cannot execute sql query: %w", err)
	}

	// extracting result from table (first and only one row)
	var cnt int
	if err := row.Scan(&cnt); err != nil {
		return 0, fmt.Errorf("cannot read table row: %w", err)
	}

	return cnt, nil
}
