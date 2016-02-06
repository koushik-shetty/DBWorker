package app

import (
	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Database struct {
	db     *sqlx.DB
	config *DBConfig
}

func (d *Database) Config() *DBConfig {
	return d.config
}

func NewDatabase(dbc *DBConfig) (*Database, error) {
	db, err := sqlx.Connect(dbc.Driver(), dbc.DBConnectionString())
	if err != nil {
		return nil, err
	}
	return &Database{
		db:     db,
		config: dbc,
	}, nil
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}
