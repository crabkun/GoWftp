package controllers

import (
	"ftpd/models"
	"os"
	"path/filepath"
	"strconv"

	"strings"
	"ftpd/conf"

	"github.com/astaxie/beego"
	"encoding/json"
)

//跳转Controller
type AdminController struct {
	beego.Controller
}

//跳转Controller
type AdminUserController struct {
	beego.Controller
}
type AdminUserCheckPathController struct {
	beego.Controller
}
type AdminUserDeleteController struct {
	beego.Controller
}
type AdminUserGetInfoController struct {
	beego.Controller
}
type AdminUserEditController struct {
	beego.Controller
}
type AdminSysController struct {
	beego.Controller
}
type AdminSysGetController struct {
	beego.Controller
}
type AdminSysSetController struct {
	beego.Controller
}
//判断是否已登陆并跳转
func (c *AdminController) Get() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	c.Redirect("/admin/user", 302)
}
func (c *AdminSysController) Get() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	c.Data["IsSys"] = true
	c.TplName = "admin.tpl"
}
func (c *AdminSysGetController) Get() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	b,_:=models.GetAllSetting()
	c.Ctx.ResponseWriter.Write(b)
	c.Ctx.ResponseWriter.WriteHeader(200)
}
func (c *AdminSysSetController) Post() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	data:=c.GetString("d")
	var tmp []models.Config
	if err:=json.Unmarshal([]byte(data),&tmp);err!=nil{
		c.StopRun()
	}
	//TODO
	if len(tmp)>0 {
		if err:=models.SetSetting(tmp[0].Name,tmp[0].Value);err!=nil{
			println(err.Error())
		}
	}
	c.StopRun()
}
func (c *AdminUserController) Get() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	var list []models.User
	models.GetAllUser(&list)
	var tmp string
	var stat string
	var size string
	var readwritedelete string
	var totalsize int64
	for _, v := range list {
		totalsize = 0
		readwritedelete = "读:x 写:x 删:x"
		switch v.Type {
		case 0:
			stat = "未激活"
		case 1:
			stat = "普通用户"
		case 2:
			stat = "管理员"
		default:
			stat = "异常"
		}
		if v.Read {
			readwritedelete = strings.Replace(readwritedelete, "读:x", "读:√", 1)
		}
		if v.Write {
			readwritedelete = strings.Replace(readwritedelete, "写:x", "写:√", 1)
		}
		if v.Delete {
			readwritedelete = strings.Replace(readwritedelete, "删:x", "删:√", 1)
		}
		if v.Path == "" {
			size = "0KB"
		} else {
			path := conf.Ftppath+ "/" + v.Path
			_, err1 := os.Stat(path)
			if err1 == nil {
				filepath.Walk(path, func(f string, info os.FileInfo, err error) error {
					totalsize += info.Size()
					return err
				})
				size = strconv.FormatInt(totalsize/1024, 10) + "KB"
			} else {
				size = "目录损坏"
			}

		}

		tmp += `<tr>
        <td>` + strconv.Itoa(v.Id) + `</td>
        <td>` + v.Username + `</td>
        <td>` + stat + `</td>
        <td>` + size + `</td>
		<td>` + readwritedelete + `</td>
        <td><btn class="btn btn-info" data-toggle="modal" data-target="#editModal" onclick="edituserconfirm(this)">修改</button></td>
        <td><btn class="btn btn-danger" data-toggle="modal" data-target="#delModal" onclick="deleteuserconfirm(this)">删除</button></td>
        </tr>`
	}
	c.Data["UserList"] = beego.Str2html(tmp)
	c.Data["IsUser"] = true
	c.TplName = "admin.tpl"
}
func (c *AdminUserCheckPathController) Post() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	id := c.GetString("id")
	if id == "" {
		c.StopRun()
	}
	idn, _ := strconv.Atoi(id)
	msg := models.GetPathUser(idn)
	if msg == "" {
		c.Ctx.WriteString("删除后此用户的文件目录也同时删除掉且无法恢复！" + msg)
		c.StopRun()
	}
	c.Ctx.WriteString("此用户与下列用户共用同一目录，此次删除不会删除掉文件目录，除非你把下列所有用户都删除：<br/>" + msg)
	c.StopRun()
}
func (c *AdminUserDeleteController) Post() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	id := c.GetString("id")
	if id == "" || id == "1" {
		c.Ctx.WriteString("操作失败！")
		c.StopRun()
	}
	idn, _ := strconv.Atoi(id)
	tmp, err1 := models.GetUserInfo(idn)
	if err1 != nil {
		c.Ctx.WriteString("操作失败！")
		c.StopRun()
	}
	msg := models.GetPathUser(idn)
	if msg == "" && tmp.Path != "" {
		os.RemoveAll(conf.Ftppath + "/" + tmp.Path)
	}
	models.DeleteUser(tmp)
	c.Ctx.WriteString("操作成功！")
	c.StopRun()
}
func (c *AdminUserGetInfoController) Post() {
	if CheckLogin(&c.Controller, 2) == false {
		c.StopRun()
	}
	id := c.GetString("id")
	idn, _ := strconv.Atoi(id)
	tmp, err := models.GetUserInfo(idn)
	if err != nil {
		c.Ctx.WriteString("操作失败！")
		c.StopRun()
	}
	msg := "|r|w|d"
	if tmp.Read {
		msg = strings.Replace(msg, "r", "1", 1)
	} else {
		msg = strings.Replace(msg, "r", "0", 1)
	}
	if tmp.Write {
		msg = strings.Replace(msg, "w", "1", 1)
	} else {
		msg = strings.Replace(msg, "w", "0", 1)
	}
	if tmp.Delete {
		msg = strings.Replace(msg, "d", "1", 1)
	} else {
		msg = strings.Replace(msg, "d", "0", 1)
	}
	msg = strconv.Itoa(tmp.Type) + "|" + tmp.Path + msg
	c.Ctx.WriteString(msg)
}
func (c *AdminUserEditController) Post() {
	ret := "<script>alert('msg');location.href='/admin/user';</script>"
	if CheckLogin(&c.Controller, 2) == false {
		c.Ctx.WriteString(strings.Replace(ret, "msg", "操作失败！权限不足", -1))
		c.StopRun()
	}
	id, usertype, password, path, canread, canwrite, candelete := c.GetString("id"), c.GetString("usertype"), c.GetString("password"), c.GetString("path"), c.GetString("read"), c.GetString("write"), c.GetString("delete")
	if id == "" || usertype == "" {
		c.Ctx.WriteString(strings.Replace(ret, "msg", "操作失败！参数不完整", -1))
		c.StopRun()
	}
	idn, _ := strconv.Atoi(id)
	tmp, err := models.GetUserInfo(idn)
	if err != nil {
		c.Ctx.WriteString(strings.Replace(ret, "msg", "操作失败！用户不存在", -1))
		c.StopRun()
	}
	if !models.CheckPathAvailable(path) {
		c.Ctx.WriteString(strings.Replace(ret, "msg", "操作失败！目录名字不符合命名规范（不得含有..,<>,/,\\,|,:,\"\",*,?）", -1))
		c.StopRun()
	}
	if idn != 1 {
		switch usertype {
		case "0":
			tmp.Type = 0
		case "1":
			tmp.Type = 1
		case "2":
			tmp.Type = 2
		}
	}
	if password != "" {
		tmp.Password = password
	}

	if canread == "1" {
		tmp.Read = true
	} else {
		tmp.Read = false
	}

	if canwrite == "1" {
		tmp.Write = true
	} else {
		tmp.Write = false
	}

	if candelete == "1" {
		tmp.Delete = true
	} else {
		tmp.Delete = false
	}

	tmp.Path = path
	pathmsg := ""
	realpath := conf.Ftppath + path
	if _, patherr := os.Stat(realpath); patherr != nil {
		if err:=os.Mkdir(realpath, 0777);err!=nil{
			println("创建目录失败，原因：",err.Error())
		}
		pathmsg = "系统检测到此目录不存在，已经成功帮你创建此目录。"
	}
	models.EditUser(tmp)
	c.Ctx.WriteString(strings.Replace(ret, "msg", "操作成功！"+pathmsg, -1))
}
