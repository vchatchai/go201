package routers

import (
	"github.com/astaxie/beego"
	"github.com/vchatchai/go201/my-first-beego-project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
}
