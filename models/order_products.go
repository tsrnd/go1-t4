package models

import (
	"github.com/astaxie/beego/orm"
)

type OrderProducts struct {
	Id       int64 `orm:"auto"`
	Quantity int
}

func init() {
	orm.RegisterModel(new(OrderProducts))
}
