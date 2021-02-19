package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "kuldeep.goyal@kreditzy.com"
	main.Data["EmailName"] = "Kuldeep Goyal"
	main.TplName = "index.tpl"
}
func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "kuldeep.goyal@kreditzy.com"
	main.Data["EmailName"] = "Kuldeep Goyal"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "default/hello-sitepoint.tpl"
}
