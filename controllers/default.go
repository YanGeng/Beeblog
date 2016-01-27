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

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}

	c.TplNames = "home.html"
}
