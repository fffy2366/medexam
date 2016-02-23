package routers

import (
	"github.com/astaxie/beego"
	"medexam/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//method 首字母需要大写
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/login", &controllers.UserController{}, "post:Dologin")
	beego.Router("/reg", &controllers.UserController{}, "get:Reg")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/forget-password", &controllers.UserController{}, "get:Forget")
	//beego.AutoRouter(&controllers.UserController{})
	beego.Include(&controllers.CMSController{})
	beego.Include(&controllers.ExamplesController{})
}
