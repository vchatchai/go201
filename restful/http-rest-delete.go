package restful

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	var route Route = Route{"deleteEmployee", "DELETE", "/employee/delete", deleteEmployee}

	routes = append(routes, route)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := new(Employee)
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		panic(err)
	}
	for id, emp := range employees {
		if employee.Id == emp.Id {
			employees = append(employees[:id], employees[id+1:]...)
			break
		}
	}

	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(employees)
}

func DeleteRest() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	router := AddRoutes(muxRouter)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}
