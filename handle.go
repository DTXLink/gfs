package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func (c *Context) server(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//key := params.Get("k")
	//callback := params.Get("cb")
	path := r.URL.Path

	if path == "/" {
		home(w, r)
	} else {
		key := path[1:len(path)]
		//fmt.Println("key:" + key)
		c.download(w, r, key)
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

func (c *Context) upload(w http.ResponseWriter, r *http.Request) {
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

	key := fmt.Sprintf("%s%s", gen_md5_str(data), ext)

	c.store.Set(key, data)
	if err != nil {
		//fmt.Println("upload file fail:" key)
		fmt.Println(err)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s", key)))
}

func (c *Context) download(w http.ResponseWriter, r *http.Request, key string) {
	data, err := c.store.Get(key)
	if err != nil {
		fmt.Fprint(w, "the file not exits!")
	}

	path := key
	request_type := path[strings.LastIndex(path, "."):]

	//fmt.Println(request_type)
	switch request_type {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	case ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	}

	//fmt.Fprint(w, data)
	//w.Header().Set("Accept-Ranges", "bytes")
	//w.Header().Set("Content-Length", len(data))
	//w.Write(data)
	io.Copy(w, bytes.NewReader(data))
}
