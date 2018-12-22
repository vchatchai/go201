package restful

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

const WEB_SERVICE_HOME string = "http://localhost:8080"

func getEmployeesClient() {
	response, err := resty.R().Get(WEB_SERVICE_HOME + "/employees")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response.Body()))

}

func postEmployeeClient() {
	response, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Employee{Id: "03", LastName: "Vichai"}).
		Post(WEB_SERVICE_HOME + ":" + "/employee/add")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(response.Body()))
}

func updateEmployeeClient() {
	reponse, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Employee{Id: "03", FirstName: "Chatchai", LastName: "Vichai"}).
		Put(WEB_SERVICE_HOME + ":" + "/employee/update")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reponse.Body()))

}

func deleteEmployeeClient() {
	response, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Employee{Id: "2"}).
		Delete(WEB_SERVICE_HOME + ":" + "/employee/delete")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response.Body()))
}
