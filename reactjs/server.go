package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"getEmployees", "GET", "/employees", getEmployees},
	Route{"addEmployee", "POST", "/employee/add", addEmployee},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName`
	LastName  string `json:lastName`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {

	// body, _ := ioutil.ReadAll(r.Body)
	// log.Println(string(body))

	newEmployee := Employee{Id: string(len(employees))}
	json.NewDecoder(r.Body).Decode(&newEmployee)

	if newEmployee.FirstName != "" {
		newEmployee.LastName = newEmployee.FirstName
		employees = append(employees, newEmployee)
	}
	json.NewEncoder(w).Encode(employees)
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	muxrouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxrouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/build/")))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server  ::", err)
	}
}
