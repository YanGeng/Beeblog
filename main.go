package main

import (
	//"Beeblog/controllers"
	"Beeblog/models"
	_ "Beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 开启orm调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 注册 beego 路由
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/login", &controllers.LoginController{})

	beego.Run()
}
