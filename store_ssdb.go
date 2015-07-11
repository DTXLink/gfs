package gfs

import (
	//"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

type SSDBStorage struct {
	client *ssdb.Client
}

func NewSSDBStorage(cfg *Config) *SSDBStorage {
	c, err := ssdb.Connect(cfg.Storage.SSDBHost, cfg.Storage.SSDBPort)
	if err != nil {
		os.Exit(1)
	}
	return &SSDBStorage{client: c}
}

func (this *SSDBStorage) set(key string, val []byte) error {
	this.client.Do("set", key, val)
	return nil
}

func (this *SSDBStorage) get(key string) (interface{}, error) {
	var val interface{}
	var err error
	val, err = this.client.Get(key)
	if err != nil {
		return nil, err
	}
	return val, err
}
