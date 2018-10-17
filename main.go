package main

import (
	_ "github.com/udistrital/minutas_argo_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default","postgres","postgres://" + beego.AppConfig.String("PGuser") + ":" + beego.AppConfig.String("PGpass") + "@" + beego.AppConfig.String("PGurls") + "/" + beego.AppConfig.String("PGdb") + "?sslmode=disable&search_path=" + beego.AppConfig.String("PGschemas") + "&timezone=UTC")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

