package main

import (
	"flag"
	"fmt"
	gofs "github.com/DTXLink/gofs"
	"runtime"
)

func main() {

	flag.Parse()
	runtime.GOMAXPROCS(500)

	fmt.Println("web start...")

	gofs.ServeHTTP()
}
