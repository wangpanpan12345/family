package routers

import (
	"family/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/neo4j", &controllers.MainController{}, "*:Test")

	beego.Router("/admin", &controllers.MainController{}, "*:Admin")

	ns :=
		beego.NewNamespace("/member",
			beego.NSRouter("/", &controllers.FamilyController{}, "get:Get"),
			beego.NSRouter("/node", &controllers.FamilyController{}, "get:AddMember"),
			beego.NSRouter("/list", &controllers.FamilyController{}, "get:ListMember"),
			beego.NSRouter("/q", &controllers.FamilyController{}, "get:SearchMember"),
			beego.NSRouter("/node/ref", &controllers.FamilyController{}, "get:ShowThreeMember"),
		)

	beego.AddNamespace(ns)

}
