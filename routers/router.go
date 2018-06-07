package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:IndexGet")

	beego.Router("/goods", &controllers.MainController{}, "*:GetResult")
}
