package html

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GorillaMux() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
