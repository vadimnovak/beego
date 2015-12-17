package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "VADNOV.COM"
	main.Data["Email"] = "vadnov2013@gmail.com"
	main.Data["EmailName"] = "VadNov"
	main.TplNames = "default/hello-sitepoint.tpl"
}