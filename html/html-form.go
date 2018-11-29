package html

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/gorilla/mux"
)

type User struct {
	Username string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		template, _ := template.ParseFiles("login-form.html")

		template.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		user := new(User)

		decoder := schema.NewDecoder()

		decodeErr := decoder.Decode(user, r.PostForm)

		if decodeErr != nil {
			log.Printf("error mapping parsed form data to struct:", decodeErr)
		}

		fmt.Fprintf(w, "User: %s Password: %s", user.Username, user.Password)
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
