package main

//import (
//	"github.com/ssdb/gossdb/ssdb"
//	"os"
//)

//type IStore interface {
//	Set(key string, data []byte) error
//	Get(key string) ([]byte, error)
//}

//type SSDBStore struct {
//	client *ssdb.Client
//}

//func NewSSDBStore(cfg *Config) *SSDBStore {
//	c, err := ssdb.Connect(cfg.Storage.SSDBHost, cfg.Storage.SSDBPort)
//	if err != nil {
//		os.Exit(1)
//	}
//	return &SSDBStore{client: c}
//}

//func (this *SSDBStore) Set(key string, val []byte) error {
//	this.client.Do("set", key, val)
//	return nil
//}

//func (this *SSDBStore) Get(key string) (interface{}, error) {
//	var val interface{}
//	var err error
//	val, err = this.client.Get(key)
//	if err != nil {
//		return nil, err
//	}
//	return val, err
//}
