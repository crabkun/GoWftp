package conf

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"os"
	"strings"
	"log"
	"strconv"
)


var Ftppath string
var Ftpport int

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func printVer() string {
	return "1.0"
}
func title() string {
	return beego.AppConfig.DefaultString("title", "多用户FTP系统")
}

func init(){
	Ftppath=beego.AppConfig.DefaultString("ftppath", getCurrentDirectory()+"/ftpfile/")
	Ftpport,_=strconv.Atoi(beego.AppConfig.DefaultString("ftpport", "21"))
	if Ftppath==""{
		log.Fatal("没有配置ftp文件夹")
	}
	beego.AddFuncMap("ver", printVer)
	beego.AddFuncMap("title", title)
}