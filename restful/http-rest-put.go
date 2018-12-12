package restful

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	routes = append(routes, Route{"updateEmployee", "PUT", "/employee/update", updateEmployee})
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		panic(err)
	}

	var isUpsert = true
	for idx, emp := range employees {
		if emp.Id == employee.Id {
			isUpsert = false
			log.Printf("updating employee id :: %s with firstName as  :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			break
		}
	}

	if isUpsert {
		log.Printf("upserting employee id :: %s  with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
		employees = append(employees, employee)
	}

	json.NewEncoder(w).Encode(employees)

}

func Update() {
	router := mux.NewRouter().StrictSlash(true)

	router = AddRoutes(router)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		panic(err)
	}
}
