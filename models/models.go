package models

import (
	"os"

	"fmt"

	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //驱动
	"encoding/json"
)

//User 用户结构
type User struct {
	Id       int
	Type     int
	Username string `orm:"size(255)"`
	Password string `orm:"size(255)"`
	Path     string `orm:"size(100)"`
	Read     bool
	Write    bool
	Delete   bool
}
//Config 配置结构
type Config struct {
	Id int
	Name string `orm:"size(255)"`
	Value string `orm:"size(255)"`
}
func init() {
	err := orm.RegisterDataBase("default", "sqlite3", "./db")
	if err != nil {
		println("[错误]数据库连接失败：", err.Error())
		os.Exit(0)
	}
	orm.RegisterModel(new(User),new(Config))
	if _, initerr := orm.NewOrm().Raw("select * from init").Exec(); initerr != nil {
		println("数据库没有初始化！现在开始初始化")
		orm.RunSyncdb("default", false, true)
		orm.NewOrm().Insert(&User{Id: 1, Type: 2, Username: "admin", Password: "admin"})
		orm.NewOrm().Insert(&Config{Name:"canreg",Value:"true"})
		orm.NewOrm().Raw("create table init(a int)").Exec()
		println("数据库初始化完成！默认管理账号admin密码admin")
	}

}

//AddUser 添加用户
func AddUser(user string, pass string, path string, read bool, write bool, delete bool) {
	o := orm.NewOrm()
	tmp := User{Type: 0, Username: user, Password: pass, Path: path, Read: read, Write: write, Delete: delete}
	o.Insert(&tmp)
}

//GetUserType 获取用户信息（可用于登陆验证）
func GetUserType(user string, pass string) (bool, int, int, string) {
	o := orm.NewOrm()
	var tmp []User
	o.QueryTable("user").Filter("username", user).Filter("password", pass).All(&tmp)
	if len(tmp) == 0 {
		return false, 0, 0, ""
	}
	return true, tmp[0].Id, tmp[0].Type, tmp[0].Path
}

//GetUserExist 用户是否存在
func GetUserExist(user string) bool {
	o := orm.NewOrm()
	num, _ := o.QueryTable("user").Filter("username", user).Count()
	if num == 0 {
		return false
	}
	return true
}

//GetAllUser 取所有用户
func GetAllUser(tmp *[]User) {
	o := orm.NewOrm()
	o.QueryTable("user").All(tmp)
}

//GetPathUser 取同一目录的所有用户
func GetPathUser(id int) string {
	o := orm.NewOrm()
	var tmp User
	err := o.QueryTable("user").Filter("id", id).One(&tmp)
	if err != nil || tmp.Path == "" {
		return ""
	}
	var tmp1 []User
	var msg string
	o.QueryTable("user").Filter("path", tmp.Path).Exclude("id", id).All(&tmp1)
	for _, v := range tmp1 {
		msg += fmt.Sprintf("用户名：%s (ID:%d)，用户目录：%s<br/>", v.Username, v.Id, v.Path)
	}
	return msg
}

//GetUserInfo 取用户
func GetUserInfo(id int) (User, error) {
	o := orm.NewOrm()
	var tmp User
	err1 := o.QueryTable("user").Filter("id", id).One(&tmp)
	if err1 != nil {
		return tmp, err1
	}
	return tmp, nil
}

//DeleteUser 删除用户
func DeleteUser(tmp User) {
	o := orm.NewOrm()
	o.Delete(&tmp)
}

//EditUser 编辑用户
func EditUser(tmp User) {
	o := orm.NewOrm()
	o.Update(&tmp)
}

//CheckPathAvailable 检查目录名字是否符合规范
func CheckPathAvailable(name string) bool {
	var unlist = []string{"<", ">", "/", "\\", "|", ":", "\"", "*", "?", ".."}
	for _, tmp := range unlist {
		if strings.Index(name, tmp) != -1 {
			return false
		}
	}
	return true
}

//GetSetting 获取设置
func GetAllSetting() ([]byte,error){
	var tmp []Config
	o := orm.NewOrm()
	o.QueryTable("config").All(&tmp)
	return json.Marshal(tmp)
}
func SetSetting(key string,value string) error{
	o := orm.NewOrm()
	var tmp Config
	o.QueryTable("config").Filter("name", key).One(&tmp)
	tmp.Value=value
	_,err:=o.Update(&tmp)
	return err
}
func GetSetting(key string) string{
	o := orm.NewOrm()
	var tmp Config
	o.QueryTable("config").Filter("name", key).One(&tmp)
	return tmp.Value
}