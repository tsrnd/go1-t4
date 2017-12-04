package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type ProductGroup struct {
	gorm.Model	
	Name    	string		`schema:"name"`
	Products	[]Product	//has many products
}

func (productGroup *ProductGroup) GetRelationship() map[string]interface{}{
	relationship := map[string]interface{} {
		"Products": &productGroup.Products,
	}
	return relationship
}

func GetProductGroups() (productGroups []ProductGroup, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Find(&productGroups).Error
	})
	return productGroups, err
}

func GetProductGroup(id uint) (productGroup ProductGroup, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("id = ?", id).Find(&productGroup).Error
	})
	return
}

func GetProductsByGroupID(id uint) (products []Product, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("group_id = ?", id).Find(&products).Error
	})
	return products, err
}
