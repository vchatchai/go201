package https

import (
	"fmt"
	"net/http"
)

const (
	CONN_HOST          = "localhost"
	CONN_PORT          = "8443"
	HTTPS_CERTIFICATE  = "domain.crt"
	DOMAIN_PRIVATE_KEY = "domain.key"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func Https() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServeTLS(CONN_HOST+":"+CONN_PORT, HTTPS_CERTIFICATE, DOMAIN_PRIVATE_KEY, nil)
	if err != nil {
		panic(err)
	}

}
