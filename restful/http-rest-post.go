package restful

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	routes = append(routes, Route{"addEmployee", "POST", "/employee/add", addEmployee})

	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Quz"},
	}

	routes = append(routes, Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	})
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		panic(err)
	}

	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)

	employees = append(employees, employee)

	json.NewEncoder(w).Encode(employees)
}

func Post() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
