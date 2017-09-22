package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"errors"
)

type UserInfo struct {
	Id       int `orm:"pk;auto"`
	Username string `orm:"size(50)" form:"username" valid:"MaxSize(50);AlphaNumeric" `
	Email    string `orm:"size(50);unique" form:"email" valid:"MaxSize(50);Required;Email" `
	Nickname string `orm:"size(50)" form:"nickname" valid:"MaxSize(50)"`
	Password string `orm:"size(50)" form:"password" valid:"MaxSize(50);Required"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
}

func (i *UserInfo) TableName() string {
	return "tl_user_info"
}

func (i *UserInfo) Valid(v *validation.Validation) {
	//v.SetError("Repassword", "两次输入的密码不一样")

	logger.Debug("now validation userInfo")
}

//验证用户信息
func (i *UserInfo) CheckUser() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&i)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (i *UserInfo) Get(uid int) {
	o := orm.NewOrm()
	i.Id = uid
	err := o.Read(i)
	if err == orm.ErrNoRows {
		logger.Error("没有数据")
	} else if err == orm.ErrMissPK {
		logger.Error("找不到主键")
	} else {
		logger.Debug("查询成功!")
	}
}

func (i *UserInfo) Insert() {
	o := orm.NewOrm()
	id, err := o.Insert(i)
	if err == nil {
		logger.Debug("insert userinfo success!")
	} else {
		logger.Error("insert userinfo failed!")
	}
	logger.Debug("id = ", id)
}

func (i *UserInfo) BatchInsert(list []*UserInfo) (int, error) {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, list)
	if err == nil {
		logger.Debug("batch insert success")
		logger.Debug("insert num: %d", successNums)
	} else {
		logger.Error("batch insert failed!")
	}
	return int(successNums), err
}

func (i *UserInfo) Update()  {
	o := orm.NewOrm()
	num, err := o.Update(i)
	//num, err := o.QueryTable(i).Filter("name", "slene").Update(orm.Params{
	//	"name": "astaxie",
	//})
	if err == nil {
		logger.Debug("update success")
		logger.Debug("%d row update", num)
	} else {
		logger.Error("update failed!")
	}
}

func (i *UserInfo) Delete(uid int)  {
	o := orm.NewOrm()
	i.Id = uid
	if num, err := o.Delete(i); err == nil {
		logger.Debug("delete success, row: %d", num)
	}
}

func (i *UserInfo) GetByEmail(email string)  {
	o := orm.NewOrm()
	// 也可以直接使用对象作为表名
	qs := o.QueryTable(i)
	//qs.Filter("id", 1) // WHERE id = 1
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("Profile__Age", 18) // 使用字段名和 Field 名都是允许的
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
	//qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
	//qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20)
	//qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)
	// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000

	var maps []orm.Params
	num, err := qs.Filter("email", email).Values(&maps)
	if err == nil {
		logger.Debug("Result Nums: %d\n", num)
		for _, m := range maps {
			logger.Debug("info:", m["Id"], m["Email"])
		}
	}
	for index, param := range maps {
		logger.Debug("line: ", index)
		logger.Debug("id", param["Id"])
	}
}

func (i *UserInfo) ListUserInfo() []*UserInfo {
	//var userInfoList []*UserInfo
	//num, err := orm.NewOrm().QueryTable(new(UserInfo)).All(&userInfoList)
	return nil
}




func init() {
	orm.RegisterModel(new(UserInfo))
}
