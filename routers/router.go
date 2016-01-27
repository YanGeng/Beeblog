package routers

import (
	"Beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
}
