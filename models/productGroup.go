package models

import (
	"github.com/goweb4/database"
)

type ProductGroup struct {
	Model
	Name     string    `schema:"name"`
	Products []Product //has many products
}

func GetProductGroups() (productGroups []ProductGroup, err error) {
	rows, err := database.DBCon.Db.Query("SELECT id, name FROM product_groups")
	if err != nil {
		return
	}
	for rows.Next() {
		productGroup := ProductGroup{}
		err = rows.Scan(&productGroup.ID, &productGroup.Name)
		if err != nil {
			return
		}
		productGroups = append(productGroups, productGroup)
	}
	rows.Close()
	return
}

// func GetProductsByGroupID(id uint) (products []Product, err error) {
// 	err = database.DBCon.Where("group_id = ?", id).Find(&products).Error
// 	for i, _ := range products {
// 		database.DBCon.Model(products[i]).Related(&products[i].Images)
// 	}
// 	return products, err
// }

func GetProductGroupByID(id uint) (*ProductGroup, error) {
	productGroup := ProductGroup{}
	err := database.DBCon.Db.
		QueryRow("SELECT id, name FROM product_groups WHERE id = $1", id).
		Scan(&productGroup.ID, &productGroup.Name)
	return &productGroup, err
}
