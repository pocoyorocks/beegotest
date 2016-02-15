package conf

import (
	"log"
	"path"
	
	"github.com/kelseyhightower/envconfig"		
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	
	"thanpogi/common/models"
	"thanpogi/common/structs"
)

const (
	APP_NAME string = "authentication"
	APP_HTTP_PORT int = 8016
)

var appenv structs.AppEnvConfig

func init() {
	err := envconfig.Process("thanpogigoapp", &appenv)
	if err == nil {
		beego.AppConfigPath = path.Join(appenv.Confighome, "app.conf")
		beego.AppConfigProvider = "ini"

		beego.BConfig.AppName = APP_NAME
		beego.BConfig.Listen.HTTPPort = APP_HTTP_PORT
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir = map[string]string{"/swagger": appenv.Swaggerhome}
				
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", appenv.Databasehome + "/orm_test.db")
		orm.RegisterModel(new(models.Article))
	}

	if err != nil {
		log.Fatal(err)
	}

}