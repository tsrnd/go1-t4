package models

import (
	"github.com/astaxie/beego/orm"
)

type Orders struct {
	Id         int64  `orm:"auto"`
	TotalMoney string `orm:"size(128)"`
	Status     string `orm:"size(128)"`
	OrderDate  string `orm:"size(128)"`
	Payment    *Payments `orm:"rel(fk)"`
	User       *Users `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Orders))
}