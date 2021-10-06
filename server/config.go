package server

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Addr string
		Port string
	}
	Database struct {
		Addr     string
		Port     string
		User     string
		Password string
		Database string
	}
}

func (c *Config) ReadConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("error occured in file %q: %v", filename, err)
	}

	return c, nil
}
