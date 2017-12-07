package models

import (
	"github.com/astaxie/beego/orm"
)

type ProductGroups struct {
	Id   int64  `orm:"auto"`
	Name string `orm:"size(128)"`
	Products []*Products `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(ProductGroups))
}
