package app

import (
	"DBWorker/lib"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Driver   string `json: "driver"`
	Hostname string `json: "hostname"`
	Dbname   string `json: "dbaname"`
	Schema   string `json: "schema"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		Driver:   "postgres",
		Hostname: "localhost",
		Dbname:   "koteldb",
		Schema:   "wall",
		Username: "pilgrim",
		Password: "western_wall",
	}
}

func (dbconf *DBConfig) DBConnectionString() string {
	return fmt.Sprintf("host=%s dbname=%s search_path=%s user=%s password=%s sslmode=disable", dbconf.Hostname, dbconf.Dbname, dbconf.Schema, dbconf.Username, dbconf.Password)
}

func (dbconf *DBConfig) DBSetupString() string {
	return fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", dbconf.Hostname, dbconf.Username, dbconf.Password)
}

func NewDBConfig(host, dbname, schema, username, password string) *DBConfig {
	return &DBConfig{
		Driver:   "postgres",
		Hostname: host,
		Dbname:   dbname,
		Schema:   schema,
		Username: username,
		Password: password,
	}
}

type Database struct {
	db     *sqlx.DB
	config *DBConfig
}

func (d *Database) Config() *DBConfig {
	return d.config
}

func NewDatabase(dbc *DBConfig, connectionString string) (*Database, *lib.Error) {
	db, err := sqlx.Connect(dbc.Driver, connectionString)
	if err != nil {
		return nil, lib.NewError(lib.DBError, "newdatabase", fmt.Sprintf("error connection to database:%v", err))
	}
	return &Database{
		db:     db,
		config: dbc,
	}, nil
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, *lib.Error) {
	result, err := d.db.Exec(query, args...)
	return result, lib.ToLibError(err, "", "")
}
