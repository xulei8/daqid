package routers

import (
	"github.com/astaxie/beego"
	"github.com/xulei8/daqid/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gocrm_login", &controllers.MainController{})
	beego.Router("/gocrm_help", &controllers.HelpController{})
	beego.Router("/gocrm_crmadmin", &controllers.CrmAdmin{})
	beego.Router("/gocrm_agent", &controllers.Agent{})
	beego.Router("/gocrm_agentcall", &controllers.AgentCall{})
	beego.Router("/gocrm_report", &controllers.MainController{})
	beego.Router("/gocrm_crmimport", &controllers.CrmImport{})
	beego.Router("/gocrm_agentcallaction", &controllers.CallAction{})
	beego.Router("/gocrm_agentcalllog", &controllers.CallLog{})

}
