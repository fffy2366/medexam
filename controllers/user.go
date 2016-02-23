package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"medexam/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	user := c.GetSession("user")
	if user != nil {
		c.Ctx.WriteString("has login")
	} else {
		c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "login.html"
	}

}
func (this *UserController) Dologin() {
	username := this.GetString("username")
	password := this.GetString("password")
	//验证登录
	user, err := models.Find(username, password)
	if !err {
		//fmt.Printf("ERR: %v\n", err)
		//this.Ctx.WriteString("username:" + username + " password:" + password)
		mystruct := map[string]interface{}{"retcode": 1, "retmsg": "no user", "nickname": user.Nickname}
		this.Data["json"] = &mystruct
		this.ServeJSON()
		return
	} else {
		//判断账号是否正确
		fmt.Printf("Pass: %v\n", password)

		pass := Md5(password)
		fmt.Printf("userPass: %v\n", user.Password)
		fmt.Printf("Pass: %v\n", pass)
		mystruct := map[string]interface{}{"": ""}
		if pass == user.Password {
			mystruct = map[string]interface{}{"retcode": 0, "retmsg": "success", "nickname": user.Nickname}
			//保存到session
			this.SetSession("user", user)
		} else {
			mystruct = map[string]interface{}{"retcode": 2, "retmsg": "vali pass"}
		}
		this.Data["json"] = &mystruct
		this.ServeJSON()
		//this.Ctx.WriteString("username:" + username + " password:" + password)
	}
	//fmt.Printf("Mobile: %v\n", user.Mobile)
	//fmt.Printf("Nickname: %v\n", user.Nickname)

}

func (this *UserController) Logout() {
	this.DelSession("user")
	mystruct := map[string]interface{}{"retcode": 0, "retmsg": "success"}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}

func (c *UserController) Reg() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "reg.html"
}
func (c *UserController) Forget() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "forget_password.html"
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 sharejs.com
	//fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}
