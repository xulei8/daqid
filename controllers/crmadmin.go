package controllers

import (
	"github.com/astaxie/beego"
)

type CrmAdmin struct {
	beego.Controller
}

func (this *CrmAdmin) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin.tpl"
}
