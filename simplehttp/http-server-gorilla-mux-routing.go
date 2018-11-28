package simplehttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Post Request"))
	})
var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		w.Write([]byte("Hi " + name))
	},
)

func Gorilla() {
	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
