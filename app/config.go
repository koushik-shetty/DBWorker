package app

import (
	"fmt"
)

type DBConfig struct {
	driver   string
	hostname string
	dbname   string
	schema   string
	username string
	password string
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

// func DefaultDBConfig() *DBConfig {
// 	return &DBConfig{
// 		driver:   "postgres",
// 		hostname: "localhost",
// 		dbname:   "koteldb",
// 		schema:   "wall",
// 		username: "pilgrim",
// 		password: "western_wall",
// 	}
// }

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		driver:   "postgres",
		hostname: "localhost",
		dbname:   "koteldb",
		schema:   "public",
		username: "postgres",
		password: "postgres",
	}
}
func (dbconf *DBConfig) DBConnectionString() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbconf.HostName(), dbconf.DBName(), dbconf.UserName(), dbconf.Password())
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
