package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
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

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// Read User by columns
func (u *User) ReadByUsername() (err error) {
	o := orm.NewOrm()
	err = o.Read(u, "UserName")
	return
}
