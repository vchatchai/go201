package html

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		template, _ := template.ParseFiles("login-form.html")

		template.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		user := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Fprintf(w, "User: %s Password: %s", user, password)
	}

}

func HtmlForm() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)

	if err != nil {
		panic(err)
	}

}
