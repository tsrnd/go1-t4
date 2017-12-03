package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type ProductGroup struct {
	gorm.Model
	Name string `schema:"name"`
}

func GetProductGroups() (productGroups []ProductGroup, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Find(&productGroups).Error
	})
	return productGroups, err
}

func GetProductsByGroupID(id uint) (products []Product, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("group_id = ?", id).Find(&products).Error
	})
	return products, err
}
