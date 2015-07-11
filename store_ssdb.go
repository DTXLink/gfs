package gfs

import (
	//"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

type SSDBStorage struct {
	db *ssdb.Client
}

func NewSSDBStorage(c *Context) *SSDBStorage {
	z := new(SSDBStorage)
	ip := c.cfg.Storage.SSDBHost
	port := c.cfg.Storage.SSDBPort
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		os.Exit(1)
	}
	z.db = db
	return z
}

func (z *SSDBStorage) save_file(key string, val []byte) error {

	z.db.Do("set", key, val)

	return nil
}

func (z *SSDBStorage) get_file(key string) (interface{}, error) {

	var val interface{}
	var err error
	val, err = z.db.Get(key)

	if err != nil {
		return nil, err
	}

	return val, err
}
