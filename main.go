package main

import (
	"flag"
	"fmt"
	//log "github.com/golang/glog"
)

func main() {
	//var err error
	flag.Parse()

	fmt.Println("web start.")

	//b64 := encode_base64("heell")
	//fmt.Println(b64)

	StartHTTP() // start http listen.

	//log.Infoln("helll")
}
