package restful

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func Version() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)

	AddRoutes(muxRouter.PathPrefix("/v1").Subrouter())
	AddRoutes(muxRouter.PathPrefix("/v2").Subrouter())

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
