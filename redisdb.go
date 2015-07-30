package main

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisDB struct {
	isConnect bool
	pool      *redis.Pool
}

func NewRedisDB(s string, p int) (*RedisDB, error) {
	addr := fmt.Sprintf("%s:%d", s, p)
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &RedisDB{
		pool:      pool,
		isConnect: true,
	}, nil
}

func (this *RedisDB) getConnect() (redis.Conn, error) {
	if this.isConnect {
		conn := this.pool.Get()
		return conn, nil
	} else {
		return nil, errors.New("Can not connect db")
	}
}

func (this *RedisDB) exist(key string) bool {
	conn, err := this.getConnect()
	if err != nil {
		return false
	}
	defer conn.Close()

	isExists, _ := redis.Bool(conn.Do("EXISTS", key))
	return isExists
}

func (this *RedisDB) get(key string) ([]byte, error) {
	conn, err := this.getConnect()
	if err != nil {
		return nil, errors.New("Can not connect db!")
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *RedisDB) set(key string, val []byte) error {
	_, err := this.Do("set", key, val)
	return err
}

func (this *RedisDB) Do(commandName string, args ...interface{}) (interface{}, error) {
	conn, err := this.getConnect()
	if err != nil {
		return nil, errors.New("Can not connect db!")
	}
	defer conn.Close()
	return conn.Do(commandName, args...)
}

func (this *RedisDB) send(commandName string, args ...interface{}) error {
	conn, err := this.getConnect()
	if err != nil {
		return errors.New("Can not connect db!")
	}
	defer conn.Close()
	return conn.Send(commandName, args...)
}

func (this *RedisDB) flush() {
	if this.isConnect {
		conn := this.pool.Get()
		defer conn.Close()
		conn.Flush()
	}
}

func (this *RedisDB) Close() {
	if this.isConnect {
		this.pool.Close()
	}
}

/*
Hashmap
*/
//		fmt.Println("helllo ssdb set")
//		c.store.hset("lin", "name", []byte("linfx"))
//		name, err := c.store.hget("lin", "name")
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(string(name))

func (this *RedisDB) hset(name string, key string, val []byte) error {
	_, err := this.Do("hset", name, key, val)
	return err
}

func (this *RedisDB) hget(name string, key string) ([]byte, error) {
	conn, err := this.getConnect()
	if err != nil {
		return nil, errors.New("Can not connect db!")
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("hget", name, key))
	if err != nil {
		return nil, err
	}
	return data, nil
}
