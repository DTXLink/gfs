package main

import (
	"bytes"
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	//"strings"
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

	w.Header().Set("content-type", "text/html; charset=utf-8")
	//w.Write([]byte(html))
	io.WriteString(w, html)
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
		//fmt.Println(err)
		http.Error(w, "Not Found", 404)
		return
	}
	defer file.Close()

	//update file
	//fmt.Println("upload file:%s", handle.Filename)
	//fmt.Println("ext" + path.Ext(handle.Filename))
	ext := path.Ext(handle.Filename)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		//ctx.doError(err, 500)
		//fmt.Println(err)
		http.Error(w, "Not Found", 404)
		return
	}

	//key := fmt.Sprintf("%s%s", gen_md5_str(data), ext)
	key := gen_md5_str(data)

	//save to ssdb
	//c.store.Set(key, data)
	c.store.hset(key, "format", []byte(ext))
	c.store.hset(key, "data", data)

	if err != nil {
		//fmt.Println("upload file fail:" key)
		//fmt.Println(err)
		http.Error(w, "Not Found", 404)
		return
	}
	ret := key
	//ret := fmt.Sprintf("http://%s/%s", r.Host, key)
	//fmt.Println(ret)
	w.Write([]byte(ret))
}

func (c *Context) download(w http.ResponseWriter, r *http.Request, key string) {
	//get file form ssdb

	format, err := c.store.hget(key, "format")
	if err != nil {
		//fmt.Fprint(w, "the file not exits!")
		http.Error(w, "Not Found", 404)
		return
	}

	//data, err := c.store.Get(key)
	data, err := c.store.hget(key, "data")
	if err != nil {
		//fmt.Fprint(w, "the file not exits!")
		http.Error(w, "Not Found", 404)
		return
	}

	//path := key
	//request_type := path[strings.LastIndex(path, "."):]

	//fmt.Println(request_type)
	switch string(format) {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	case ".jpg":
		w.Header().Set("content-type", "image/jpeg")
	case ".amr":
		w.Header().Set("content-type", "audio/amr")
	default:
		w.Header().Set("content-type", "application/octet-stream")
	}
	//fmt.Fprint(w, data)
	//w.Header().Set("Accept-Ranges", "bytes")
	//w.Header().Set("Content-Length", len(data))
	//w.Write(data)
	io.Copy(w, bytes.NewReader(data))
}
