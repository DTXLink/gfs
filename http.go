package main

import (
	"fmt"
	"log"
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

	redisdb, err := NewRedisDB(cfg.Store.SSDBHost, cfg.Store.SSDBPort)
	if err != nil {
		return nil, err
	}

	return &Context{
		config: &cfg,
		store:  redisdb,
	}, nil
}

func StartHTTP(c *Context) {
	//http.HandleFunc("/", c.server)
	//http.HandleFunc("/upload", c.upload)

	// internal
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("/", c.server)
	httpServeMux.HandleFunc("/upload", c.upload)

	//http listen
	bind := fmt.Sprintf("%s:%d", c.config.System.Host, c.config.System.Port)
	//http.ListenAndServe(addr, nil)
	fmt.Printf("http listen addr:\"%s\"", bind)
	go httpListen(httpServeMux, bind)
}

func httpListen(mux *http.ServeMux, bind string) {
	server := &http.Server{Handler: mux, Addr: bind}
	server.SetKeepAlivesEnabled(false)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error(%v)", err)
		panic(err)
	}
}
