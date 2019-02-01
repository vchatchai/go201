package controllers

import (
	"github.com/astaxie/beego"
)

type FirstController struct {
	beego.Controller
}
type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: 1, FirstName: "Foo", LastName: "Bar"},
		Employee{Id: 2, FirstName: "Baz", LastName: "Quz"},
	}
}

func (this *FirstController) GetEmployees() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = employees
	this.ServeJSON()
}

func (this *FirstController) Dashboard() {
	this.Data["employees"] = employees
	this.TplName = "dashboard.tpl"
}
