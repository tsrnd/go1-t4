package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int64  `orm:"auto"`
	UserName string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	Password string `orm:"type(longtext)"`
	Gender   string `orm:"size(128)"`
	Role     string `orm:"size(128)"`
	Avatar   string `orm:"size(128)"`
	Phone    string `orm:"size(128)"`
	Provider string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int64) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.QueryTable(new(Users)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}