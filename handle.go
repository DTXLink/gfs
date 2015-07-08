package gofs

import (
	//log "code.google.com/p/log4go"
	//myrpc "github.com/Terry-Mao/gopush-cluster/rpc"
	"net/http"
	//"strconv"
	//"time"
	"fmt"
	"io/ioutil"
)

func doDefault(w http.ResponseWriter, r *http.Request) {

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

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	//w.Write([]byte(`handle hello world`))

	//z := getConnect()
	//val, err := get_file("key1")
	//	if err != nil {
	//	}
	//fmt.Printf("%s\n", val)

	//w.Write([]byte(val))

	doDefault(w, r)
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world v1`))
	save_file("key1", []byte(`hello i mmm`))
}

func upLoad(w http.ResponseWriter, r *http.Request) {
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

	//md5Sum, err := z.storage.SaveImage(data)
	save_file("key1", data)
	if err != nil {
		//z.doError(err, 500)
		return
	}

	fmt.Fprint(w, fmt.Sprintf("upload success! md5 : %s", data))
}
