package controllers

import (
	"strconv"

	"sitepointgoapp/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

type ManageController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "hello.tpl"
}

func (manage *ManageController) Delete() {
	// convert the string value to an int
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")
	article := models.Article{}

	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
		if num, err := o.Delete(&models.Article{Id: articleId}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
}
