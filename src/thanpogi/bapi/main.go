package main

import (
	_ "thanpogi/bapi/docs"
	_ "thanpogi/bapi/conf"
	_ "thanpogi/bapi/routers"
		
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
