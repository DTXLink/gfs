package gfs

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (z *ZContext) server(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//key := params.Get("k")
	//callback := params.Get("cb")
	//protoStr := params.Get("p")
	path := r.URL.Path

	if path == "/" {
		z.home(w, r)
	} else {
		md5key := path[1:len(path)]
		fmt.Println("md5key:" + md5key)
		z.download(w, r, md5key)
	}
}

func (z *ZContext) home(w http.ResponseWriter, r *http.Request) {
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

func (c *ZContext) upload(w http.ResponseWriter, r *http.Request) {
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

	md5key := fmt.Sprintf("%s", gen_md5_str(data))

	c.storage.save_file(md5key, data)
	if err != nil {
		//fmt.Println("upload file fail:" md5key)
		return
	}

	//fmt.Println("upload file: %s" md5key)
	w.Write([]byte(fmt.Sprintf("%s", md5key)))
}

func (c *ZContext) download(w http.ResponseWriter, r *http.Request, key string) {

	//	format := strings.Trim(r.Form.Get("f"), " ")
	//	if len(format) == 0 {
	//		//f = "none"
	//		//f = imgInfo.Format
	//		format = ".amr"
	//	}

	val, err := c.storage.get_file(key)
	if err != nil {
		fmt.Fprint(w, "the file not exits!")
	}
	//file, err := os.Open(val)
	fmt.Fprint(w, val)
}
