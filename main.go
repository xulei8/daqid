/*
daqi crm Go项目。
xulei8@qq.com
GPL.
Jan 18, 2014
项目使用bee创建
port 808
*/
package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xulei8/daqid/routers"
)

func main() {
	beego.Run()
}
