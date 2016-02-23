package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Users struct {
	Id       int `orm:"pk;column(user_id);"`
	Nickname string
	Mobile   string
	Password string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8", 30)
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users))
}

func Find(username string, password string) (Users, bool) {
	//user := Users{Mobile: username, Password: password}
	//user := Users{Id: 1}
	user := Users{Mobile: username}

	o := orm.NewOrm()
	err := o.Read(&user, "Mobile")
	if err != nil {
		return user, false
	} else {

		return user, true
	}
}
