package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("p", "8000", "port")

func main() {
	flag.Parse()
	fmt.Println("Starting server on http://0.0.0.0:" + *port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+*port, nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "index.html")
	return
}
