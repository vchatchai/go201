package html

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Println("Error occurred while executing the template or writing its output: ", err)
		return
	}
}

func FirstTemplate() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
