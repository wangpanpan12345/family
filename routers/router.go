package routers

import (
	"family/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/neo4j", &controllers.MainController{},"*:Test")
    beego.Router("/admin", &controllers.MainController{},"*:Admin")
}
