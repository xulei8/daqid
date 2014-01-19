package controllers

import (
	"github.com/astaxie/beego"
)

type Agent struct {
	beego.Controller
}

func (this *Agent) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

type AgentCall struct {
	beego.Controller
}

func (this *AgentCall) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
