package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Products struct {
	Id      int64  `orm:"auto"`
	Name    string `orm:"size(128)"`
	Size    string `orm:"size(128)"`
	Color   string `orm:"size(128)"`
	Price   string `orm:"size(128)"`
	InStock int
	ProductGroups *ProductGroups `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Products))
}

// AddProducts
func AddProducts(m *Products) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductsById retrieves Products by Id. Returns error if
// Id doesn't exist
func GetProductsById(id int64) (v *Products, err error) {
	o := orm.NewOrm()
	v = &Products{Id: id}
	if err = o.QueryTable(new(Products)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProducts 
func GetAllProducts() (products []Products, err error) {
	// 
	return products, err
}

// DeleteProducts
func DeleteProducts(id int64) (err error) {
	o := orm.NewOrm()
	v := Products{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Products{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
