package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"] = append(beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"],
		beego.ControllerComments{
			"Delete",
			`/delete/:id([0-9]+)`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"] = append(beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"],
		beego.ControllerComments{
			"Update",
			`/update/:id([0-9]+)`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"] = append(beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"],
		beego.ControllerComments{
			"View",
			`/view`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"] = append(beego.GlobalControllerRouter["thanpogi/bapi/controllers:ManageController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

}
