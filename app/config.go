package app

import (
	"encoding/json"
	"io/ioutil"

	"DBWorker/lib"
)

type Config struct {
	dbconfig *DBConfig
}

type configJSON struct {
	DbConfig DBConfig `json:"database"`
}

func DefaultConfig() *Config {
	return &Config{
		dbconfig: DefaultDBConfig(),
	}
}

func (c *Config) DBConfig() *DBConfig {
	return c.dbconfig
}

func LoadConfig(filename string) (*Config, *lib.Error) {
	if filename == "" {
		return DefaultConfig(), nil
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, lib.ToLibError(err, lib.FileError, "LoadConfig")
	}

	c := &configJSON{}
	err = json.Unmarshal(file, c)
	if err != nil {
		return nil, lib.ToLibError(err, lib.JSONError, "LoadConfig")
	}
	return jsonToConfig(c), nil

}

func jsonToConfig(c *configJSON) *Config {
	return &Config{
		dbconfig: &c.DbConfig,
	}
}
