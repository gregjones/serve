package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
	dir  string
)

func init() {
	flag.IntVar(&port, "port", 8080, "Port to run server on.")
	flag.StringVar(&dir, "dir", ".", "Directory to serve.")
	flag.Parse()
}

func main() {
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
}
