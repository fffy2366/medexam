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
	//beego 升级后TplNames变为TplName
	//c.TplName = "index.tpl"
	c.TplName = "index.html"
}

func (c *MainController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	// c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *MainController) StaticBlock() {
	this.TplName = "index.tpl"
}
