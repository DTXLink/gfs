package gfs

import (
	//log "code.google.com/p/log4go"
	//"encoding/json"
	//"fmt"
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

func StartHTTP(z *Context) {

	http.HandleFunc("/", z.server)
	http.HandleFunc("/upload", z.upload)

	http.ListenAndServe(":3000", nil)
}
