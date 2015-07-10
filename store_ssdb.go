package gfs

import (
	//"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

type ZSSDBStorage struct {
	context *ZContext
	db      *ssdb.Client
}

func NewSSDBStorage(c *ZContext) *ZSSDBStorage {
	z := new(ZSSDBStorage)
	ip := "192.168.82.2"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		os.Exit(1)
	}
	z.db = db
	z.context = c
	return z
}

func (z *ZSSDBStorage) save_file(key string, val []byte) error {

	z.db.Do("set", key, val)

	return nil
}

func (z *ZSSDBStorage) get_file(key string) (interface{}, error) {

	var val interface{}
	var err error
	val, err = z.db.Get(key)

	if err != nil {
		return nil, err
	}

	return val, err
}
