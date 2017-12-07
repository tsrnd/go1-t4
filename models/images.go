package models

import (
	"github.com/astaxie/beego/orm"
)

type Images struct {
	Id   int64  `orm:"auto"`
	Name string `orm:"size(128)"`
	Url  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Images))
}