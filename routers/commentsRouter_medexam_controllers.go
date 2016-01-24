package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"AllBlock",
			`/all/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"Info",
			`/info.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:MainController"] = append(beego.GlobalControllerRouter["medexam/controllers:MainController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

}
