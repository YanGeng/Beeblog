package controllers

import (
	"Beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}

	c.TplName = "home.html"

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories
}
