package simplehttp

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}

func HttpServer() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
