package controllers

import (
	"github.com/astaxie/beego"
	"family/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {

	test := models.Member{}

	test.QuickNDirty()
	
}

func (c *MainController) Admin() {
	c.TplName = "admin.tpl"
}