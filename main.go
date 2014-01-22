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
	m "github.com/xulei8/daqid/models"
	_ "github.com/xulei8/daqid/routers"
	"os"
)

func main() {
	initArgs()
	beego.Run()

}

//建立数据库
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "--newdb" || v == "new" || v == "newdb" {
			m.Syncdb()
			os.Exit(0)

		}
	}
}
