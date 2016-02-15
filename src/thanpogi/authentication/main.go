package main

import (
	_ "thanpogi/authentication/docs"
	_ "thanpogi/authentication/conf"
	_ "thanpogi/authentication/routers"
		
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
