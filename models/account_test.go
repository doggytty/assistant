package models

import (
	"testing"
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	pc, file, line, ok := runtime.Caller(1)
	logger.Debug("%v,%v,%v,%v", pc, file, line, ok)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	logger.Debug(apppath)
	beego.TestBeegoInit(apppath)

	SyncDataBase()
}

func TestSyncDataBase(t *testing.T) {
	// 读取beego的配置
	SyncDataBase()
}

func TestUserInfo_Insert(t *testing.T) {
	info := new(UserInfo)
	info.Email = "doggytty@126.com"
	info.Nickname = "doggytty"
	info.Password = "tpc"
	info.Username = "lucien"
	info.Insert()
	logger.Debug("%d", info.Id)
}

func TestUserInfo_BatchInsert(t *testing.T) {
	info := new(UserInfo)
	info.Email = "doggytty@126.com"
	info.Nickname = "doggytty"
	info.Password = "tpc"
	info.Username = "lucien"

	info1 := new(UserInfo)
	info1.Email = "doggytty@136.com"
	info1.Nickname = "doggytty"
	info1.Password = "tpc"
	info1.Username = "lucien"

	list := make([]*UserInfo, 2)
	list[0] = info
	list[1] = info1

	num, err := info.BatchInsert(list)
	logger.Debug("%v, %v", num, err)

	time.Sleep(time.Second*1)
}

func TestUserInfo_Delete(t *testing.T) {
	info := new(UserInfo)
	info.Delete(1)
}

func TestUserInfo_GetByEmail(t *testing.T) {

}

func TestUserInfo_Update(t *testing.T) {

}

func TestUserInfo_Get(t *testing.T) {

}
