package main

import (
	_ "ftpd/routers"
	"ftpd/server"
	"ftpd/driver"
	"ftpd/auth"
	"github.com/astaxie/beego"
	_ "ftpd/conf"
	"ftpd/conf"
	"os"
)

func main() {
	factory := &driver.FileDriverFactory{}
	opts := &server.ServerOpts{
		Factory: factory,
		Auth:    &auth.SimpleAuth{},
		Port:    conf.Ftpport,
		Log:     false,
	}
	server := server.NewServer(opts)
	go func(){
		if err:=server.ListenAndServe();err!=nil{
			println("FTP服务器启动失败！原因：",err.Error())
			os.Exit(-1)
		}
	}()
	beego.Run()
}
