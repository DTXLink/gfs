package gfs

import (
	//"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

type SSDBStorage struct {
	db *ssdb.Client
}

func NewSSDBStorage(ctx *Context) *SSDBStorage {
	client := new(SSDBStorage)
	ip := ctx.cfg.Storage.SSDBHost
	port := ctx.cfg.Storage.SSDBPort
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		os.Exit(1)
	}
	client.db = db
	return client
}

func (client *SSDBStorage) save_file(key string, val []byte) error {
	client.db.Do("set", key, val)
	return nil
}

func (client *SSDBStorage) get_file(key string) (interface{}, error) {
	var val interface{}
	var err error
	val, err = client.db.Get(key)
	if err != nil {
		return nil, err
	}
	return val, err
}
