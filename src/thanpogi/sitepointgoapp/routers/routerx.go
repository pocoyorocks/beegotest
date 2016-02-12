package routers

import (
	"github.com/astaxie/beego"
	"thanpogi/sitepointgoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/manage/add", &controllers.ManageController{}, "get,post:Add")
	beego.Router("/manage/view", &controllers.ManageController{}, "get:View")
	beego.Router("/manage/home", &controllers.ManageController{}, "*:Home")
	beego.Router("/manage/delete/:id([0-9]+)", &controllers.ManageController{}, "*:Delete")
	beego.Router("/manage/update/:id([0-9]+)", &controllers.ManageController{}, "*:Update")
}
