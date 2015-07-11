package gfs

import (
	//log "code.google.com/p/log4go"
	//"encoding/json"
	"fmt"
	//"net"
	"net/http"
	//"time"
)

type Context struct {
	cfg     *Config
	storage *SSDBStorage
}

func NewContext(cfgFile string) (*Context, error) {

	cfg, err := LoadConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	c := new(Context)
	c.cfg = &cfg
	c.storage = NewSSDBStorage(c)
	return c, nil
}

func StartHTTP(ctx *Context) {

	http.HandleFunc("/", ctx.server)
	http.HandleFunc("/upload", ctx.upload)

	addr := fmt.Sprintf("%s:%d", ctx.cfg.System.Host, ctx.cfg.System.Port)
	http.ListenAndServe(addr, nil)
}
