package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPtr := flag.String("config", "", "config file")
	flag.Usage = usage
	flag.Parse()

	if *configPtr == "" {
		*configPtr = "./conf/config.conf"
	}

	cfgFile := *configPtr
	isExist, _ := exists(cfgFile)
	if !isExist {
		fmt.Println("config file not exist!")
		os.Exit(-1)
	}

	context, err := NewContext(cfgFile)

	if err != nil {
		panic(err)
		os.Exit(-2)
	}

	fmt.Println("web start..")

	StartHTTP(context)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:--config=/etc/config.ini \n")
	flag.PrintDefaults()
	os.Exit(-2)
}
