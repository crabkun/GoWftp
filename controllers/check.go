package controllers

import "github.com/astaxie/beego"

//跳转Controller
type CheckController struct {
	beego.Controller
}

//判断是否已登陆并跳转
func (c *CheckController) Get() {
	if c.GetSession("username") == nil {
		c.Redirect("/login", 302)
	} else {
		switch c.GetSession("usertype") {
		case 0:
			c.Ctx.WriteString("<script>alert('帐号尚未激活，请等待管理员激活后再试.');location.href='/logout';</script>")
		case 1:
			c.Redirect("/index", 302)
		case 2:
			c.Redirect("/admin", 302)
		default:
			c.Ctx.WriteString("<script>alert('帐号状态异常，请联系管理员.');location.href='/logout';</script>")
		}

	}

}

//CheckLogin 检测登录状态，供其他函数调用
func CheckLogin(c *beego.Controller, needtype int) bool {
	if c.GetSession("usertype") != needtype {
		c.Ctx.WriteString("<script>alert('帐号状态异常，请联系管理员.');location.href='/logout';</script>")
		return false
	}
	return true
}
