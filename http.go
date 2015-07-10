package main

import (
	//log "code.google.com/p/log4go"
	//"encoding/json"
	//"fmt"
	//"net"
	"net/http"
	//"time"
)

func StartHTTP() {
	//httpServeMux := http.NewServeMux() // external
	//httpServeMux.HandleFunc("/v1/server/get", GetServer)

	http.HandleFunc("/", get)
	http.HandleFunc("/get", getInfo)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":3000", nil)
}
