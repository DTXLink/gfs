package main

import (
	"flag"
	"fmt"
	"os"
	//log "github.com/golang/glog"
	gfs "github.com/DTXLink/gfs"
)

func main() {
	//var err error
	flag.Parse()

	fmt.Println("web start.")

	//b64 := encode_base64("heell")
	//fmt.Println(b64)

	zContext, err := gfs.NewContext()

	if err != nil {
		panic(err)
		os.Exit(-2)
	}

	gfs.StartHTTP(zContext)
}
