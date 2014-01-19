package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/daqid/models"
	"github.com/xulei8/daqid/utils"
	"strings"
)

var o orm.Ormer

type baseRouter struct {
	beego.Controller

	isLogin bool
}

type CCT struct {
	ID int
	NN string
}
type CrmAdmin struct {
	//beego.Controller
	object m.DqContact
	baseRouter
}

func (this *CCT) Test() (link string) {

	return "testcct"
}
func (this *baseRouter) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["Paginator"] = p
	this.Data["Tests"] = "aaaaaaaaaaaa"
	cct := CCT{ID: 2, NN: "ssss"}
	this.Data["Cc"] = cct
	return p
}

type CrmImport struct {
	beego.Controller
}

func (this *CrmAdmin) Object() interface{} {
	return &this.object
}

func (this *CrmAdmin) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	var articles []m.DqContact

	o := orm.NewOrm()

	qs := o.QueryTable("dq_contact").Limit(10, 20).OrderBy("-id")
	//fmt.Println("%q",qs)
	this.SetPaginator(10, 33)
	if err := this.SetObjects(qs, &articles); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}

	this.TplNames = "admin.tpl"
}

func (this *CrmAdmin) SetObjects(qs orm.QuerySeter, objects interface{}) error {
	_, err := qs.Count()
	//fmt.Println("Count %d",cnt)
	if err != nil {
		return err
	}
	// create paginator

	if cnt, err := qs.RelatedSel().All(objects); err != nil {
		return err
	} else {
		this.Data["Objects"] = objects
		this.Data["ObjectsCnt"] = cnt

	}
	return nil
}

func (this *CrmImport) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "import.tpl"
}

// 导入数据
func (this *CrmImport) Post() {
	nums := this.GetString("nums")
	line := strings.Split(nums, "\n")
	l := len(line)
	i := 0

	o = orm.NewOrm()
	for i = 0; i < l; i++ {

		a := strings.Trim(line[i], "\n")
		b := strings.Trim(a, " ")
		if len(b) > 3 {
			u := new(m.DqContact)
			u.Mname = b
			u.Uname = b
			u.Tel = b
			//fmt.Println("line " + ":" + b)
			//m.AddDqContact(u)
			//o = orm.NewOrm()

			o.Insert(u)
		}
	}

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "import.tpl"
}
