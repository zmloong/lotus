package ftp

import (
	"io"
	"time"

	"github.com/jlaffaye/ftp"
)

type (
	FileEntry struct {
		FileName string
		FileSize uint64
		ModTime  int64
		IsDir    bool
	}
	IFtp interface {
		Append(path string, r io.Reader) (err error)
		ChangeDir(path string) (err error)
		ChangeDirToParent() (err error)
		Delete(path string) (err error)
		Login(user string, password string) (err error)
		Logout() (err error)
		MakeDir(path string) (err error)
		NoOp() (err error)
		RemoveDir(path string) (err error)
		RemoveDirRecur(path string) (err error)
		Rename(from string, to string) (err error)
		SetTime(path string, t time.Time) (err error)
		Stor(path string, r io.Reader) (err error)
		StorFrom(path string, r io.Reader, offset uint64) (err error)
		Type(transferType ftp.TransferType) (err error)
		IsGetTimeSupported() (rst bool)
		IsSetTimeSupported() (rst bool)
		IsTimePreciseInList(path string) (rst bool)
		Walk(root string) *ftp.Walker
		CurrentDir() (str string, err error)
		FileSize(path string) (size int64, err error)
		GetEntry(path string) (file *FileEntry, err error)
		GetTime(path string) (t time.Time, err error)
		List(path string) (files []*FileEntry, err error)
		NameList(path string) (entries []string, err error)
		Retr(path string) (frs *ftp.Response, err error)
		RetrFrom(path string, offset uint64) (frs *ftp.Response, err error)
	}
)

var defsys IFtp

func OnInit(config map[string]interface{}, option ...Optionfn) (err error) {
	defsys, err = newSys(newOptions(config, option...))
	return
}
func NewSys(option ...Optionfn) (sys IFtp, err error) {
	sys, err = newSys(newOptionsByOptionFn(option...))
	return
}

func Append(path string, r io.Reader) (err error) {
	return defsys.Append(path, r)
}
func ChangeDir(path string) (err error) {
	return defsys.ChangeDir(path)
}
func ChangeDirToParent() (err error) {
	return defsys.ChangeDirToParent()
}
func Delete(path string) (err error) {
	return defsys.Delete(path)
}
func Login(user string, password string) (err error) {
	return defsys.Login(user, password)
}
func Logout() (err error) {
	return defsys.Logout()
}
func MakeDir(path string) (err error) {
	return defsys.MakeDir(path)
}
func NoOp() (err error) {
	return defsys.NoOp()
}
func RemoveDir(path string) (err error) {
	return defsys.RemoveDir(path)
}
func RemoveDirRecur(path string) (err error) {
	return defsys.RemoveDirRecur(path)
}
func Rename(from string, to string) (err error) {
	return defsys.Rename(from, to)
}
func SetTime(path string, t time.Time) (err error) {
	return defsys.SetTime(path, t)
}
func Stor(path string, r io.Reader) (err error) {
	return defsys.Stor(path, r)
}
func StorFrom(path string, r io.Reader, offset uint64) (err error) {
	return defsys.StorFrom(path, r, offset)
}
func Type(transferType ftp.TransferType) (err error) {
	return defsys.Type(transferType)
}
func IsGetTimeSupported() (rst bool) {
	return defsys.IsGetTimeSupported()
}
func IsSetTimeSupported() (rst bool) {
	return defsys.IsSetTimeSupported()
}
func IsTimePreciseInList(path string) (rst bool) {
	return defsys.IsTimePreciseInList(path)
}
func Walk(root string) *ftp.Walker {
	return defsys.Walk(root)
}
func CurrentDir() (str string, err error) {
	str, err = defsys.CurrentDir()
	return
}
func FileSize(path string) (size int64, err error) {
	size, err = defsys.FileSize(path)
	return
}
func GetEntry(path string) (entry *FileEntry, err error) {
	entry, err = defsys.GetEntry(path)
	return
}
func GetTime(path string) (t time.Time, err error) {
	t, err = defsys.GetTime(path)
	return
}
func List(path string) (entries []*FileEntry, err error) {
	entries, err = defsys.List(path)
	return
}
func NameList(path string) (entries []string, err error) {
	entries, err = defsys.NameList(path)
	return
}
func Retr(path string) (frs *ftp.Response, err error) {
	frs, err = defsys.Retr(path)
	return
}
func RetrFrom(path string, offset uint64) (frs *ftp.Response, err error) {
	frs, err = defsys.RetrFrom(path, offset)
	return
}
