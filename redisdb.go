package main

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisDB struct {
	pool      *redis.Pool
	isConnect bool
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

func (this *RedisDB) Exist(key string) bool {
	conn, err := this.getConnect()
	if err != nil {
		return false
	}
	defer conn.Close()

	isExists, _ := redis.Bool(conn.Do("EXISTS", key))
	return isExists
}

func (this *RedisDB) Get(key string) ([]byte, error) {
	conn, err := this.getConnect()
	if err != nil {
		return nil, errors.New("Can not connect db!")
	}
	defer conn.Close()

	ff, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	fmt.Println("get:", ff)

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *RedisDB) Set(key string, val []byte) error {
	this.Do("set", key, val)
	return nil
}

func (this *RedisDB) Do(commandName string, args ...interface{}) (interface{}, error) {
	conn, err := this.getConnect()
	if err != nil {
		return nil, errors.New("Can not connect db!")
	}
	defer conn.Close()
	return conn.Do(commandName, args...)
}

func (this *RedisDB) Send(commandName string, args ...interface{}) error {
	conn, err := this.getConnect()
	if err != nil {
		return errors.New("Can not connect db!")
	}
	defer conn.Close()
	return conn.Send(commandName, args...)
}

func (this *RedisDB) Flush() {
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
