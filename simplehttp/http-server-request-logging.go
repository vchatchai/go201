package simplehttp

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Log() {
	router := mux.NewRouter()
	router.Handle("/", handlers.LoggingHandler(os.Stdout, GetRequestHandler)).Methods("GET")
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error start http server :", err)
	}
	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("GET")
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, PathVariableHandler)).Methods("GET")
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
