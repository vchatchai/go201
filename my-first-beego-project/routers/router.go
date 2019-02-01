package routers

import (
	"github.com/astaxie/beego"
	"github.com/vchatchai/go201/my-first-beego-project/controllers"
	"github.com/vchatchai/go201/my-first-beego-project/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
	beego.Router("/home", &controllers.SessionController{}, "get:Home")
	beego.Router("/login", &controllers.SessionController{}, "get:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)
}
