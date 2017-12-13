package models

import (
	"fmt"
	"github.com/goweb4/database"
)

type ProductGroup struct {
	Model
	Name     string    `schema:"name"`
	Products []Product //has many products
}

func GetProductGroups() (productGroups []ProductGroup) {
	rows, err := database.DBCon.Db.Query("SELECT id, name FROM product_groups")
	if err != nil {
		return
	}
	for rows.Next() {
		productGroup := ProductGroup{}
		err = rows.Scan(&productGroup.ID ,&productGroup.Name)
		if err != nil {
			return
		}
		productGroups = append(productGroups, productGroup)
	}
	rows.Close()
	// fmt.Println(productGroups[1].ID)
	return productGroups
}

// Get products by group
func GetProductsByGroupID(id uint) (products []Product, err error) {
	query := `SELECT id, name, size, price, in_stock, color FROM products WHERE group_id = $1`
	rows, err := database.DBCon.Db.Query(query, id)
		for rows.Next() {
			tmp := Product{}
			err = rows.Scan(&tmp.ID, &tmp.Name, &tmp.Size, &tmp.Price, &tmp.InStock, &tmp.Color)
			if err != nil {
				return
			}
			LoadImages(&tmp)
			products = append(products, tmp)
		}
		rows.Close()
	return products, err
}

// Load image from relationship one product to many images
func LoadImages(product *Product) {
	query := `SELECT url FROM images WHERE product_id = $1`
	rows, err := database.DBCon.Db.Query(query, product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		image := Image{}
		err = rows.Scan(&image.URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		product.Images = append(product.Images, image)
	}
	rows.Close()
}
