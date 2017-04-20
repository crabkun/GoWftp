package auth

import (
	"ftpd/server"
	"os"
	"ftpd/models"
	"ftpd/conf"
)

type SimpleAuth struct {
	server.Auth
}

func (*SimpleAuth) CheckPasswd(driverx server.Driver, username string, password string) (bool, error) {
	flag, id, _, path := models.GetUserType(username, password)
	realpath := conf.Ftppath + path
	_, err := os.Stat(realpath)
	if !flag || path == "" || err != nil {
		return false, nil
	}
	tmp, _ := models.GetUserInfo(id)
	driverx.SetPerm(tmp.Read, tmp.Write, tmp.Delete)
	driverx.SetRootPath(realpath)
	return true, nil
}