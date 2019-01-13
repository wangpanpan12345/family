package controllers

import (
	"family/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.LoginCheck()

	c.Redirect("/member", 301)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}

func (c *MainController) Test() {

	test := models.Member{}

	test.QuickNDirty()

}

func (c *MainController) Admin() {
	c.TplName = "admin.tpl"
}

func (c *MainController) LoginCheck() {
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

func (this *MainController) Login() {
	this.TplName = "login.tpl"
}

func (this *MainController) DCLogin() {
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	if username == "admin" && password == "_dc@2018" {
		sess := this.StartSession()
		sess.Set("finanuser", username)
		// fmt.Println(sess.Get("finanuser"))
		f := FamilyController{}
		f.Get()
		// this.S
		// this.TplName = "urlManage.tpl"
	} else {
		this.Redirect("/admin/login", 301)
	}

}
