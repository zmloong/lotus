package ftp

import (
	"fmt"
	"io"
	"time"

	"github.com/jlaffaye/ftp"
)

type Ftp struct {
	options Option
	client  *ftp.ServerConn
}

func newSys(options Option) (sys *Ftp, err error) {
	sys = &Ftp{options: options}
	err = sys.init()
	return
}

func (z *Ftp) init() (err error) {
	var (
		conn *ftp.ServerConn
	)
	if conn, err = ftp.Dial(fmt.Sprintf("%s:%d", z.options.ServerIp, z.options.Port), ftp.DialWithTimeout(time.Duration(z.options.TimeOut)*time.Second)); err != nil {
		err = fmt.Errorf("ftp.Dial err:%v", err)
		return
	}
	if err = conn.Login(z.options.UserName, z.options.PassWord); err != nil {
		err = fmt.Errorf("ftp.Login err:%v", err)
	}
	z.client = conn
	return
}
func (z *Ftp) Append(path string, r io.Reader) (err error) {
	err = z.client.Append(path, r)
	return
}
func (z *Ftp) ChangeDir(path string) (err error) {
	err = z.client.ChangeDir(path)
	return
}
func (z *Ftp) ChangeDirToParent() (err error) {
	err = z.client.ChangeDirToParent()
	return
}
func (z *Ftp) Delete(path string) (err error) {
	err = z.client.Delete(path)
	return
}
func (z *Ftp) Login(user string, password string) (err error) {
	err = z.client.Login(user, password)
	return
}
func (z *Ftp) Logout() (err error) {
	err = z.client.Logout()
	return
}
func (z *Ftp) MakeDir(path string) (err error) {
	err = z.client.MakeDir(path)
	return
}
func (z *Ftp) NoOp() (err error) {
	err = z.client.NoOp()
	return
}
func (z *Ftp) RemoveDir(path string) (err error) {
	err = z.client.RemoveDir(path)
	return
}
func (z *Ftp) RemoveDirRecur(path string) (err error) {
	err = z.client.RemoveDirRecur(path)
	return
}
func (z *Ftp) Rename(from string, to string) (err error) {
	err = z.client.Rename(from, to)
	return
}
func (z *Ftp) SetTime(path string, t time.Time) (err error) {
	err = z.client.SetTime(path, t)
	return
}
func (z *Ftp) Stor(path string, r io.Reader) (err error) {
	err = z.client.Stor(path, r)
	return
}
func (z *Ftp) StorFrom(path string, r io.Reader, offset uint64) (err error) {
	err = z.client.StorFrom(path, r, offset)
	return
}
func (z *Ftp) Type(transferType ftp.TransferType) (err error) {
	err = z.client.Type(transferType)
	return
}
func (z *Ftp) IsGetTimeSupported() (rst bool) {
	rst = z.client.IsGetTimeSupported()
	return
}
func (z *Ftp) IsSetTimeSupported() (rst bool) {
	rst = z.client.IsSetTimeSupported()
	return
}
func (z *Ftp) IsTimePreciseInList(path string) (rst bool) {
	rst = z.client.IsTimePreciseInList()
	return
}
func (z *Ftp) Walk(root string) *ftp.Walker {
	return z.client.Walk(root)
}
func (z *Ftp) CurrentDir() (str string, err error) {
	str, err = z.client.CurrentDir()
	return
}
func (z *Ftp) FileSize(path string) (size int64, err error) {
	size, err = z.client.FileSize(path)
	return
}
func (z *Ftp) GetEntry(path string) (file *FileEntry, err error) {
	var (
		entry *ftp.Entry
	)
	entry, err = z.client.GetEntry(path)
	file = &FileEntry{
		FileName: entry.Name,
		FileSize: entry.Size,
		ModTime:  entry.Time.Unix(),
		IsDir:    false,
	}
	if entry.Type == ftp.EntryTypeFolder {
		file.IsDir = true
	}
	return
}
func (z *Ftp) GetTime(path string) (t time.Time, err error) {
	t, err = z.client.GetTime(path)
	return
}
func (z *Ftp) List(path string) (files []*FileEntry, err error) {
	var (
		entries []*ftp.Entry
	)
	if entries, err = z.client.List(path); err == nil {
		files = make([]*FileEntry, len(entries))
		for i, v := range entries {
			files[i] = &FileEntry{
				FileName: v.Name,
				FileSize: v.Size,
				ModTime:  v.Time.Unix(),
				IsDir:    false,
			}
			if v.Type == ftp.EntryTypeFolder {
				files[i].IsDir = true
			}
		}
	}
	return
}
func (z *Ftp) NameList(path string) (entries []string, err error) {
	entries, err = z.client.NameList(path)
	return
}
func (z *Ftp) Retr(path string) (frs *ftp.Response, err error) {
	frs, err = z.client.Retr(path)
	return
}
func (z *Ftp) RetrFrom(path string, offset uint64) (frs *ftp.Response, err error) {
	frs, err = z.client.RetrFrom(path, offset)
	return
}
