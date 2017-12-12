package models

import (
	"github.com/goweb4/database"
)

type ProductGroup struct {
	Model
	Name     string    `schema:"name"`
	Products []Product //has many products
}

func GetProductGroups() (productGroups []ProductGroup) {
	rows, err := database.DBCon.Db.Query("SELECT name FROM product_groups")
	if err != nil {
		return
	}
	for rows.Next() {
		productGroup := ProductGroup{}
		err = rows.Scan(&productGroup.Name)
		if err != nil {
			return
		}
		productGroups = append(productGroups, productGroup)
	}
	rows.Close()
	return productGroups
}

// func GetProductsByGroupID(id uint) (products []Product, err error) {
// 	err = database.DBCon.Where("group_id = ?", id).Find(&products).Error
// 	for i, _ := range products {
// 		database.DBCon.Model(products[i]).Related(&products[i].Images)
// 	}
// 	return products, err
// }
