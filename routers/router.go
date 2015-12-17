package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
	"beego/controllers"
	"beego/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
}
