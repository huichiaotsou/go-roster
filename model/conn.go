package model

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlx(dbConfig *config.Database) (*sqlx.DB, error) {
	// Create a database connection string
	dbinfo := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbConfig.Host, dbConfig.Port, dbConfig.Username,
		dbConfig.Password, dbConfig.Name)

	// Connect to the database
	sqlx, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		return nil, fmt.Errorf("error while connecting the database: %s", err)
	}
	defer sqlx.Close()

	// Test the connection
	err = sqlx.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while pinging database: %s", err)
	}

	return sqlx, nil
}
