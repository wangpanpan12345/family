package controllers

import (
	"family/models"
	"github.com/astaxie/beego"
	"strconv"
)

type FamilyController struct {
	beego.Controller
}

func (c *FamilyController) Get() {

	c.LoginCheck()

	c.TplName = "view.tpl"
}

func (c *FamilyController) LoginCheck() {
	// sess := this.StartSession()
	sess_uid := c.GetSession("userid")
	// sess_username := this.GetSession("username")
	// beego.Debug(sess_uid, sess_username)
	beego.Debug(sess_uid)
	if sess_uid == nil {
		c.Ctx.Redirect(302, "/admin/login")
		return
	}
}

func (this *FamilyController) DCLogin() {
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	if username == "sxjp" && password == "2019" {
		sess := this.StartSession()
		sess.Set("userid", username)
		// fmt.Println(sess.Get("finanuser"))
		this.Redirect("/member", 301)
		// this.S
		// this.TplName = "urlManage.tpl"
	} else {
		this.Redirect("/admin/login", 301)
	}

}

//添加一个成员
func (c *FamilyController) AddMember() {

	s := c.Input().Get("sex")
	name := c.Input().Get("name")
	father := c.Input().Get("father")

	si, err := strconv.Atoi(s)

	sex := ""
	if si == 1 {
		sex = "son"
	} else if si == 2 {
		sex = "daughter"
	} else {
		return
	}
	if name == "" || father == "" {
		return
	}

	beego.Debug(si, sex, name, father)
	// return

	memberModel := models.Member{}
	id := memberModel.IdGen()

	beego.Debug(id)
	// return

	member := map[string]interface{}{
		"id":    id,
		"name":  name,
		"sex":   si,
		"edu":   "",
		"birth": "",
		"bf":    24,
		"info":  "",
	}
	nodeNum, refNum, err := memberModel.NewMember(member, father, []string{"father", sex})

	if err != nil {
		beego.Debug(err)
	}

	rep := Response{}
	if nodeNum == 1 && refNum == 1 {
		rep.Data = 0
		rep.Errcode = 0
		rep.Errmsg = "success"

	}
	beego.Debug(nodeNum, refNum)

	c.Data["json"] = rep
	c.ServeJSON()

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
	// id = "eb3cb4bc37d487a808bd54885a7f5e8a"
	memberModel := models.Member{}
	ms, js, diff := memberModel.ListG3Member(id)
	for k, v := range ms {
		beego.Debug(k, v)
	}

	// beego.Debug(ms)
	resp := Response{}
	resp.Data = js
	resp.Pager = ms
	resp.Errcode = 0
	resp.Diff = diff
	resp.Errmsg = "success"
	c.Data["json"] = resp
	beego.Debug(resp.Diff)
	c.ServeJSON()
	// beego.Debug(ms)
}

//搜索一个成员
func (c *FamilyController) ComputeMember() {
	id1 := c.Input().Get("id1")
	id2 := c.Input().Get("id2")

	// id = "eb3cb4bc37d487a808bd54885a7f5e8a"
	memberModel := models.Member{}
	ms, js, diff := memberModel.ComputeMember(id1, id2)
	for k, v := range ms {
		beego.Debug(k, v)
	}

	// beego.Debug(ms)
	resp := Response{}
	resp.Data = js
	resp.Pager = ms
	resp.Errcode = 0
	resp.Diff = diff
	resp.Errmsg = "success"
	c.Data["json"] = resp

	c.ServeJSON()
	// beego.Debug(ms)
}

func (c *FamilyController) ListMember() {
	memberModel := models.Member{}
	memberModel.ListMember()

}
