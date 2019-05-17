package controllers

import(
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func(c *UserController) Get() {
	c.Data["Email"] = "phandinhman@gmail.com"
	c.TplName = "user/index.tpl"
}
