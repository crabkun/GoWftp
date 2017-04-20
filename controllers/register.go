package controllers

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"ftpd/conf"
	"ftpd/models"

	"github.com/astaxie/beego"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//RegisterController 注册界面
type RegisterController struct {
	beego.Controller
}

//Get 注册界面
func (c *RegisterController) Get() {
	if c.GetSession("username") != nil {
		c.Redirect("/index", 302)
		c.StopRun()
	}
	c.Data["icon"] = "register"
	c.Data["title"] = "注册"
	c.TplName = "register.tpl"
}

//Post 注册
func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	if models.GetSetting("canreg")!="true"{
		c.Ctx.WriteString("<script>alert('管理员设置了不允许注册.');location.href='/';</script>")
		c.StopRun()
	}
	if username == "" || password == "" || repassword == "" {
		c.Ctx.WriteString("<script>alert('信息不能为空，请检查后再重试.');location.href='/register';</script>")
		c.StopRun()
	}
	if password != repassword {
		c.Ctx.WriteString("<script>alert('两次密码不一致，请检查后再重试入.');location.href='/register';</script>")
		c.StopRun()
	}
	if models.GetUserExist(username) {
		c.Ctx.WriteString("<script>alert('此用户名已经存在，请检查后再重试入.');location.href='/register';</script>")
		c.StopRun()
	}
	models.AddUser(username, password, "", false, false, false)
	_, id, _, _ := models.GetUserType(username, password)
	tmp, _ := models.GetUserInfo(id)
	tmp.Path = strconv.Itoa(id)
	models.EditUser(tmp)
	os.RemoveAll(conf.Ftppath + "\\" + strconv.Itoa(id))
	os.MkdirAll(conf.Ftppath+"\\"+strconv.Itoa(id), 0777)
	c.Ctx.WriteString("<script>alert('注册成功！');location.href='/login';</script>")
}
