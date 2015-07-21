package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (ctx *Context) server(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//key := params.Get("k")
	//callback := params.Get("cb")
	path := r.URL.Path

	if path == "/" {
		home(w, r)
	} else {
		md5key := path[1:len(path)]
		fmt.Println("md5key:" + md5key)

		ctx.store.Get("key")

		val, err := ctx.store.Get(md5key)
		if err != nil {
			fmt.Fprint(w, "the file not exits!")
		}

		fmt.Fprint(w, val)
		ctx.download(w, r, md5key)
	}
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
	            <input name="file" type="file" />
	            <input type="submit"></input>
	        </form>
	    </body>
	</html>`

	fmt.Fprint(w, html)
}

func (ctx *Context) upload(w http.ResponseWriter, r *http.Request) {
	//	if err := r.ParseMultipartForm(CACHE_MAX_SIZE); err != nil {
	//		//ctx.context.Logger.Error(err.Error())
	//		//ctx.doError(err, http.StatusForbidden)
	//		return
	//	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		//ctx.doError(err, 500)
		fmt.Println(err)
		return
	}
	defer file.Close()

	//fmt.Println("upload file:%s", handle.Filename)
	//fmt.Println("ext" + path.Ext(handle.Filename))
	ext := path.Ext(handle.Filename)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		//ctx.doError(err, 500)
		fmt.Println(err)
		return
	}

	md5key := fmt.Sprintf("%s%s", gen_md5_str(data), ext)

	ctx.store.Set(md5key, data)
	if err != nil {
		//fmt.Println("upload file fail:" md5key)
		fmt.Println(err)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s", md5key)))
}

func (ctx *Context) download(w http.ResponseWriter, r *http.Request, key string) {
	val, err := ctx.store.Get(key)
	if err != nil {
		fmt.Fprint(w, "the file not exits!")
	}
	fmt.Fprint(w, val)
}
