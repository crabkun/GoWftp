package routers

import (
	"ftpd/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.CheckController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/admin", &controllers.AdminController{})
	//beego.Router("/admin/sysinfo", &controllers.RegisterController{})
	beego.Router("/admin/user", &controllers.AdminUserController{})
	beego.Router("/admin/user/checkpath", &controllers.AdminUserCheckPathController{})
	beego.Router("/admin/user/delete", &controllers.AdminUserDeleteController{})
	beego.Router("/admin/user/edit", &controllers.AdminUserEditController{})
	beego.Router("/admin/user/getinfo", &controllers.AdminUserGetInfoController{})
	beego.Router("/admin/sys", &controllers.AdminSysController{})
	beego.Router("/admin/sys/get", &controllers.AdminSysGetController{})
	beego.Router("/admin/sys/set", &controllers.AdminSysSetController{})
}
