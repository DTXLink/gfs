package gfs

import (
	_ "fmt"
	"net/http"
)

type ZContext struct {
	storage *ZSSDBStorage
}

func NewContext() (*ZContext, error) {
	c := new(ZContext)
	c.storage = NewSSDBStorage(c)
	return c, nil
}

func StartHTTP(z *ZContext) {

	http.HandleFunc("/", z.server)
	http.HandleFunc("/upload", z.upload)

	http.ListenAndServe(":3000", nil)
}
