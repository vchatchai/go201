package simplehttp

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page")
}

func Routing() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/logout", logout)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		panic(err)
	}
}
