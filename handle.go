package main

import (
	"net/http"
	//"strconv"
	//"time"
	"fmt"
	//log "code.google.com/p/log4go"
	//myrpc "github.com/Terry-Mao/gopush-cluster/rpc"
	//"strconv"
	//"time"
	//"github.com/DTXLink/gfs/store"
	"io/ioutil"
)

func get(w http.ResponseWriter, r *http.Request) {

	//params := r.URL.Query()
	//key := params.Get("k")
	//callback := params.Get("cb")
	//protoStr := params.Get("p")
	path := r.URL.Path

	if path == "/" {
		home(w, r)
	} else if path == "/get" {
		getInfo(w, r)
	} else {
		md5key := path[1:len(path)]
		fmt.Println("md5key:" + md5key)

		val, err := get_file(md5key)
		if err != nil {
			fmt.Fprint(w, "the file not exits!")
		}

		fmt.Fprint(w, val)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	html := `<!DOCTYPE html>
	<html>
	    <head>
	        <meta charset="UTF-8"/>
	    </head>
	    <body>
	        <form action="/upload" method="POST" enctype="multipart/form-data">
	            <label for="field1">file:</label>
	            <input name="upload_file" type="file" />
	            <input type="submit"></input>
	        </form>
	    </body>
	</html>`

	fmt.Fprint(w, html)
}

// GetServer handle for server get
//func GetServer(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "GET" {
//		http.Error(w, "Method Not Allowed", 405)
//		return
//	}
//	params := r.URL.Query()
//	key := params.Get("k")
//	callback := params.Get("cb")
//	protoStr := params.Get("p")
//	res := map[string]interface{}{"ret": OK}
//	defer retWrite(w, r, res, callback, time.Now())
//	if key == "" {
//		res["ret"] = ParamErr
//		return
//	}
//	// Match a push-server with the value computed through ketama algorithm
//	node := myrpc.GetComet(key)
//	if node == nil {
//		res["ret"] = NotFoundServer
//		return
//	}
//	addrs, ret := getProtoAddr(node, protoStr)
//	if ret != OK {
//		res["ret"] = ret
//		return
//	}
//	res["data"] = map[string]interface{}{"server": addrs[0]}
//	return
//}

func upload(w http.ResponseWriter, r *http.Request) {
	//	if err := r.ParseMultipartForm(CACHE_MAX_SIZE); err != nil {
	//		//z.context.Logger.Error(err.Error())
	//		//z.doError(err, http.StatusForbidden)
	//		return
	//	}

	file, _, err := r.FormFile("upload_file")
	if err != nil {
		//z.doError(err, 500)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		//z.doError(err, 500)
		return
	}

	md5key := fmt.Sprintf("%s.wav", gen_md5_str(data))

	save_file(md5key, data)
	if err != nil {
		//fmt.Println("upload file fail:" md5key)
		return
	}

	//fmt.Println("upload file: %s" md5key)
	w.Write([]byte(fmt.Sprintf("%s", md5key)))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	//	params := r.URL.Query()
	//	callback := params.Get("cb")
	//	res := map[string]interface{}{"ret": OK}
	//	now := time.Now()
	//	defer retWrite(w, r, res, callback, now)
	//	res["data"] = map[string]interface{}{"timeid": now.UnixNano() / 100}
	//	return
}
