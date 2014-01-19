package routers

import (
	"github.com/xulei8/daqid/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
