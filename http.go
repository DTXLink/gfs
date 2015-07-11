package gfs

import (
	"fmt"
	"net/http"
)

type Context struct {
	config  *Config
	storage *SSDBStorage
}

func NewContext(cfgFile string) (*Context, error) {
	cfg, err := LoadConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return &Context{
		config:  &cfg,
		storage: NewSSDBStorage(&cfg),
	}, nil
}

func StartHTTP(ctx *Context) {

	http.HandleFunc("/", ctx.server)
	http.HandleFunc("/upload", ctx.upload)

	addr := fmt.Sprintf("%s:%d", ctx.config.System.Host, ctx.config.System.Port)
	http.ListenAndServe(addr, nil)
}
