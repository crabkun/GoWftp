package controllers

import (
	"ftpd/models"

	"github.com/astaxie/beego"
)

//LoginController 登陆界面
type LoginController struct {
	beego.Controller
}
type LogoutController struct {
	beego.Controller
}

//Get 登陆界面GET
func (c *LoginController) Get() {
	if c.GetSession("username") != nil {
		c.Redirect("/index", 302)
		c.StopRun()
	}
	c.Data["icon"] = "key"
	c.Data["title"] = "登陆"
	c.TplName = "login.tpl"
}

//Post 登陆界面POST
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	flag, id, usertype, _ := models.GetUserType(username, password)
	if !flag {
		c.Data["icon"] = "error"
		c.Data["title"] = "登陆失败，请检查用户名和密码是否错误！"
	} else {
		c.SetSession("id", id)
		c.SetSession("username", username)
		c.SetSession("usertype", usertype)
		c.Redirect("/", 302)
		c.StopRun()
	}
	c.TplName = "login.tpl"
}

//Get 注销
func (c *LogoutController) Get() {
	c.DestroySession()
	c.Redirect("/login", 302)
}
