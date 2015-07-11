package gfs

import (
	_ "errors"
	"github.com/DTXLink/gcfg"
	_ "strings"
)

type Config struct {
	System struct {
		Host string
		Port int
	}

	Storage struct {
		SSDBHost string
		SSDBPort int
	}
}

func LoadConfig(cfgFile string) (Config, error) {
	var err error
	var cfg Config

	err = gcfg.ReadFileInto(&cfg, cfgFile)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
