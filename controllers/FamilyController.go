package controllers

import (
	"family/models"
	"github.com/astaxie/beego"
)

type FamilyController struct {
	beego.Controller
}

func (c *FamilyController) Get() {

	c.TplName = "view.tpl"
}

//添加一个成员
func (c *FamilyController) AddMember() {

	memberModel := models.Member{}
	id := memberModel.IdGen()

	beego.Debug(id)
	// return

	member := map[string]interface{}{
		"id":    id,
		"name":  "王嘉硕",
		"sex":   1,
		"edu":   "",
		"birth": "",
		"bf":    24,
		"info":  "",
	}
	nodeNum, refNum, err := memberModel.NewMember(member, "98ffa4797d40cb5bac35504b9307ce37", []string{"father", "daughter"})

	if err != nil {
		beego.Debug(err)
	}

	beego.Debug(nodeNum, refNum)
	// beego.Debug(member)

}

//更新一个成员
func (c *FamilyController) UpdateMember() {

}

//删除一个成员
func (c *FamilyController) DelMember() {

}

//搜索一个成员
func (c *FamilyController) SearchMember() {
	name := c.Input().Get("name")
	memberModel := models.Member{}
	ms := memberModel.GetMember(name)

	resp := Response{}
	resp.Data = ms
	resp.Errcode = 0
	resp.Errmsg = "success"
	c.Data["json"] = resp
	c.ServeJSON()
	// beego.Debug(ms)
}

//搜索一个成员
func (c *FamilyController) ShowThreeMember() {
	id := c.Input().Get("id")
	id = "eb3cb4bc37d487a808bd54885a7f5e8a"
	memberModel := models.Member{}
	ms := memberModel.ListG3Member(id)

	beego.Debug(ms)
	resp := Response{}
	resp.Data = ms
	resp.Errcode = 0
	resp.Errmsg = "success"
	c.Data["json"] = resp
	c.ServeJSON()
	// beego.Debug(ms)
}

func (c *FamilyController) ListMember() {
	memberModel := models.Member{}
	memberModel.ListMember()

}
