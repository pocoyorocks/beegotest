package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @Title Get
// @Description Test
// @Success 200 {object} models.Object
// @Failure 403 error
// @router / [get]
func (main *MainController) Get() {
	main.Data["Website"] = "beego.me"
	main.Data["Email"] = "astaxie@gmail.com"
	main.TplName = "index.tpl"
}
