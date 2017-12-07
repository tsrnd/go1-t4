package models

import (
	"github.com/astaxie/beego/orm"
)

type Payments struct {
	Id   int64  `orm:"auto"`
	Method string `orm:"size(128)"`
	Orders []*Orders `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Payments))
}
