package simplehttp

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func GZipCompression() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))

	if err != nil {
		panic(err)
	}
}
