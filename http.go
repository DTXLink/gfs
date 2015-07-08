package gofs

import (
	//"encoding/json"
	//"fmt"
	//"net"
	"net/http"
	//"time"
	//"github.com/DTXLink/gofs/store"
)

type zhttp struct {
	//store *zssss
}

// StartHTTP start listen http.
func ServeHTTP() {
	//httpServeMux := http.NewServeMux()

	//1.0
	//httpServeMux.HandleFunc("/", home)

	//httpListen(httpServeMux, nil)

	http.HandleFunc("/", home)
	http.HandleFunc("/v1", getName)
	http.HandleFunc("/upload", upLoad)
	http.ListenAndServe(":3000", nil)

	return
}
