package main

import (
	_ "mydaily/routers"

	"github.com/astaxie/beego"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func init()  {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mydaily?charset=utf8")
	//orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetLogger("file", `{"filename":"logs/dev.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)

	beego.Run()
}
