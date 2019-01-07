package controllers

import (
	"github.com/astaxie/beego"
)

type FamilyController struct {
	beego.Controller
}

func (c *FamilyController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//添加一个成员
func (c *FamilyController) AddMember() {
	
}


//更新一个成员
func (c *FamilyController) UpdateMember() {
	
}

//删除一个成员
func (c *FamilyController) DelMember() {
	
}

//搜索一个成员
func (c *FamilyController) SearchMember() {
	
}

