package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	// c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *MainController) StaticBlock() {
	this.TplName = "index.tpl"
}
