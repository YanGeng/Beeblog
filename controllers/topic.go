package controllers

import (
	"Beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

	this.TplName = "topic.html"
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic_add.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect(("/login"), 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.UpdateTopic(tid, title, category, content)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	tid := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect(("/login"), 302)
		return
	}

	err := models.DeleteTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}

func (this *TopicController) Test() {
	// Test API: JSON string output
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "11"}
	this.ServeJSON()
	return
}
