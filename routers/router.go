package routers

import (
	"github.com/astaxie/beego"
	"github.com/najidroid/quizOfStudents/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
