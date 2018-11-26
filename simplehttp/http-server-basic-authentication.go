package simplehttp

import (
	"crypto/subtle"
	"net/http"
)

const (
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok ||
			subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 ||
			subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			rw.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			rw.WriteHeader(401)
			rw.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(rw, r)

	}
}

func HttpServerBasicAuthentication() {
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and passwrod"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		panic(err)
	}
}
