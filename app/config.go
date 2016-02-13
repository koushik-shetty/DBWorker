package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"DBWorker/lib"
)

type Config struct {
	dbconfig *DBConfig
}

type configJSON struct {
	dbConfig DBConfig
}

func LoadConfig(file string) (*Config, *lib.Error) {
	file, err := ioutil.ReadAll()
	if err != nil {
		return lib.ToLibError(err, lib.FileError, "LoadConfig")
	}
c:
	&configJSON{}
	err = json.Unmarshal(file, c)
	if err != nil {
		return lib.ToLibError(err, lib.JSONError)
	}
	return jsonToConfig(c)

}

func jsonToConfig(c *configJSON) *Config {
	return &Config{
		dbconfig: &c.dbConfig,
	}
}
