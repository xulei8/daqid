package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/daqid/models"
)

type Agent struct {
	beego.Controller
}

type CallAction struct {
	beego.Controller
}
type CallLog struct {
	beego.Controller
}

func (this *Agent) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *CallAction) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "callaction.tpl"
}

func (this *CallLog) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "calllog.tpl"
}

type AgentCall struct {
	beego.Controller
}

func (this *AgentCall) Get() {

	id, _ := this.GetInt("cid")
	if id < 1 {
		return
	}

	u := m.DqContact{Id: id}
	o := orm.NewOrm()
	o.Read(&u)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["C"] = u
	this.Data["Cid"] = id

	this.TplNames = "agentcall.tpl"
}
