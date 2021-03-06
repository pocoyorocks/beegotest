// @APIVersion 1.0.0
// @Title Than Test Beego API
// @Description Test test ets
// @Contact pocoyorocks2@xxx.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"thanpogi/bapi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/manage",
			beego.NSInclude(
				&controllers.ManageController{},
			),
		),
	)
	beego.AddNamespace(ns)
	
}
