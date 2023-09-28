package main

import (
	"flag"
	"fmt"
)

func main() {
	server := flag.Bool("s", false, "Server mode")
	flag.Parse()
	if *server {
		fmt.Println("Starting server mode ...")
		RunServer()
	} else {
		fmt.Println("Starting client mode ...")
		RunClient()
	}
}
