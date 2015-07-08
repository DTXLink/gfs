package gofs

import (
	//"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

type zstore struct {
	db *ssdb.Client
}

func (z *zstore) getConnect() {
	ip := "192.168.82.2"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		os.Exit(1)
	}
	z.db = db
}

func save_file(key string, val []byte) error {

	ip := "192.168.82.2"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		return err
	}

	db.Do("set", key, val)

	return nil
}

func get_file(key string) (interface{}, error) {

	ip := "192.168.82.2"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		return nil, err
	}

	var val interface{}
	val, err = db.Get(key)

	return val, err
}
