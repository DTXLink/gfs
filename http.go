package main

import (
	"fmt"
	"net/http"
)

type Context struct {
	config *Config
	store  *RedisDB
}

func NewContext(cfgFile string) (*Context, error) {
	cfg, err := LoadConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	redisdb, err := NewRedisDB(cfg.Storage.SSDBHost, cfg.Storage.SSDBPort)
	if err != nil {
		return nil, err
	}

	return &Context{
		config: &cfg,
		store:  redisdb,
	}, nil
}

func StartHTTP(ctx *Context) {

	http.HandleFunc("/", ctx.server)
	http.HandleFunc("/upload", ctx.upload)

	addr := fmt.Sprintf("%s:%d", ctx.config.System.Host, ctx.config.System.Port)
	http.ListenAndServe(addr, nil)
}
