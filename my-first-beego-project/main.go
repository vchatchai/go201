package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/vchatchai/go201/my-first-beego-project/controllers"
	_ "github.com/vchatchai/go201/my-first-beego-project/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
