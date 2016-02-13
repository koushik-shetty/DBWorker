package app

import (
	"DBWorker/lib"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	driver   string `json: "driver"`
	hostname string `json: "hostname"`
	dbname   string `json: "dbaname"`
	schema   string `json: "schema"`
	username string `json: "username"`
	password string `json: "password"`
}

func (dbConf *DBConfig) Driver() string {
	return dbConf.driver
}

func (dbconf *DBConfig) HostName() string {
	return dbconf.hostname
}

func (dbconf *DBConfig) DBName() string {
	return dbconf.dbname
}

func (dbconf *DBConfig) Schema() string {
	return dbconf.schema
}

func (dbconf *DBConfig) UserName() string {
	return dbconf.username
}

func (dbconf *DBConfig) Password() string {
	return dbconf.password
}

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		driver:   "postgres",
		hostname: "localhost",
		dbname:   "koteldb",
		schema:   "wall",
		username: "pilgrim",
		password: "western_wall",
	}
}

func (dbconf *DBConfig) DBConnectionString() string {
	return fmt.Sprintf("host=%s dbname=%s search_path=%s user=%s password=%s sslmode=disable", dbconf.HostName(), dbconf.DBName(), dbconf.Schema(), dbconf.UserName(), dbconf.Password())
}

func NewDBConfig(host, dbname, schema, username, password string) *DBConfig {
	return &DBConfig{
		driver:   "postgres",
		hostname: host,
		dbname:   dbname,
		schema:   schema,
		username: username,
		password: password,
	}
}

type Database struct {
	db     *sqlx.DB
	config *DBConfig
}

func (d *Database) Config() *DBConfig {
	return d.config
}

func NewDatabase(dbc *DBConfig) (*Database, *lib.Error) {
	db, err := sqlx.Connect(dbc.Driver(), dbc.DBConnectionString())
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

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.db.Exec(query, args...)
	return result, lib.ToLibError(err, "", "")
}
