package routers

import (
	"Beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.TopicController{})
	//自动路由只能匹配到/<controller>/<method>...，仅有/<controller>，自动路由如法匹配到
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.ReplyController{})
	beego.Router("/reply", &controllers.ReplyController{})
	//beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	//beego.Router("/reply/delete", &controllers.ReplyController{}, "post:Add")
}
