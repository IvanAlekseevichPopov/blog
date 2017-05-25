package routers

import (
	"sitepointgoapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/manage/delete/:id([0-9]+)", &controllers.ManageController{}, "get:Delete")
}
