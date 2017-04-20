package driver

import (
	"errors"
	"fmt"
	"ftpd/server"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileDriver struct {
	RootPath string
	Read     bool
	Write    bool
	Delete   bool
}

type FileInfo struct {
	os.FileInfo
}

func (driver *FileDriver) realPath(path string) string {
	paths := strings.Split(path, "/")
	return filepath.Join(append([]string{driver.RootPath}, paths...)...)
}

func (driver *FileDriver) Init(conn *server.Conn) {

}
func (driver *FileDriver) SetRootPath(path string) {
	driver.RootPath = path
}
func (driver *FileDriver) SetPerm(read bool, write bool, delete bool) {
	driver.Read, driver.Write, driver.Delete = read, write, delete
}
func (driver *FileDriver) ChangeDir(path string) error {
	if !driver.Read {
		return errors.New("Permission denied")
	}
	rPath := driver.realPath(path)
	f, err := os.Lstat(rPath)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	return errors.New("Not a directory")
}

func (driver *FileDriver) Stat(path string) (server.FileInfo, error) {
	if !driver.Read {
		return nil, errors.New("Permission denied")
	}
	basepath := driver.realPath(path)
	rPath, err := filepath.Abs(basepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Lstat(rPath)
	if err != nil {
		return nil, err
	}
	return &FileInfo{f}, nil
}

func (driver *FileDriver) ListDir(path string, callback func(server.FileInfo) error) error {
	if !driver.Read {
		return errors.New("Permission denied")
	}
	basepath := driver.realPath(path)
	filepath.Walk(basepath, func(f string, info os.FileInfo, err error) error {
		rPath, _ := filepath.Rel(basepath, f)
		if rPath == info.Name() {
			err = callback(&FileInfo{info})
			if err != nil {
				return err
			}
			if info.IsDir() {
				return filepath.SkipDir
			}
		}
		return nil
	})

	return nil
}

func (driver *FileDriver) DeleteDir(path string) error {
	if !driver.Delete {
		return errors.New("Permission denied")
	}
	rPath := driver.realPath(path)
	f, err := os.Lstat(rPath)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return os.Remove(rPath)
	}
	return errors.New("Not a directory")
}

func (driver *FileDriver) DeleteFile(path string) error {
	if !driver.Delete {
		return errors.New("Permission denied")
	}
	rPath := driver.realPath(path)
	f, err := os.Lstat(rPath)
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return os.Remove(rPath)
	}
	return errors.New("Not a file")
}

func (driver *FileDriver) Rename(fromPath string, toPath string) error {
	if !driver.Write {
		return errors.New("Permission denied")
	}
	oldPath := driver.realPath(fromPath)
	newPath := driver.realPath(toPath)
	return os.Rename(oldPath, newPath)
}

func (driver *FileDriver) MakeDir(path string) error {
	if !driver.Write {
		return errors.New("Permission denied")
	}
	rPath := driver.realPath(path)
	return os.Mkdir(rPath, os.ModePerm)
}

func (driver *FileDriver) GetFile(path string, offset int64) (int64, io.ReadCloser, error) {
	if !driver.Read {
		return 0, nil, errors.New("Permission denied")
	}
	rPath := driver.realPath(path)
	f, err := os.Open(rPath)
	if err != nil {
		return 0, nil, err
	}

	info, err := f.Stat()
	if err != nil {
		return 0, nil, err
	}

	f.Seek(offset, os.SEEK_SET)

	return info.Size(), f, nil
}

func (driver *FileDriver) PutFile(destPath string, data io.Reader, appendData bool) (int64, error) {
	if !driver.Write {
		return 0, errors.New("Permission denied")
	}
	rPath := driver.realPath(destPath)
	var isExist bool
	f, err := os.Lstat(rPath)
	if err == nil {
		isExist = true
		if f.IsDir() {
			return 0, errors.New("A dir has the same name")
		}
	} else {
		if os.IsNotExist(err) {
			isExist = false
		} else {
			return 0, errors.New(fmt.Sprintln("Put File error:", err))
		}
	}

	if appendData && !isExist {
		appendData = false
	}

	if !appendData {
		if isExist {
			err = os.Remove(rPath)
			if err != nil {
				return 0, err
			}
		}
		f, err := os.Create(rPath)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		bytes, err := io.Copy(f, data)
		if err != nil {
			return 0, err
		}
		return bytes, nil
	}

	of, err := os.OpenFile(rPath, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return 0, err
	}
	defer of.Close()

	_, err = of.Seek(0, os.SEEK_END)
	if err != nil {
		return 0, err
	}

	bytes, err := io.Copy(of, data)
	if err != nil {
		return 0, err
	}

	return bytes, nil
}

type FileDriverFactory struct {
	RootPath string
}

func (factory *FileDriverFactory) NewDriver() (server.Driver, error) {
	return &FileDriver{factory.RootPath, false, false, false}, nil
}
