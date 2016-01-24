package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "login.html"
}

func (c *UserController) Reg() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "reg.html"
}
