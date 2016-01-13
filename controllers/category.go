package controllers

import (
	// "fmt"
	"Beeblog/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	// "net/url"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsCategory"] = true

	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")

		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)

		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	this.TplNames = "category.html"

	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Post() {
	// this.Ctx.WriteString(fmt.Sprint(this.Input()))
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	this.Redirect("/", 301)
	return
}
