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

// Get products by group
func GetProductsByGroupID(id uint) (products []Product, err error) {
	pl := make(map[int]Product)
	query := `SELECT id, name, size, price, in_stock, color FROM products WHERE group_id = $1`
	rows, err := database.DBCon.Db.Query(query, id)
		for rows.Next() {
			product := Product{}
			err = rows.Scan(&product.ID, &product.Name, &product.Size, &product.Price, &product.InStock, &product.Color)
			if err != nil {
				fmt.Println(err)
				return
			}
			pl[int(product.ID)] = product
		}
	defer rows.Close()

	rows, err = database.DBCon.Db.Query(`SELECT url, product_id FROM images`)
	if err != nil {
		return
	}
	for rows.Next() {
		i := Image{}
		rows.Scan(&i.URL, &i.ProductId)
		product := pl[int(i.ProductId)]
		product.Images = append(pl[int(i.ProductId)].Images, i)
		pl[int(i.ProductId)] = product
	}
 	defer rows.Close()

	for _, p := range pl {
		products = append(products, p)
	}

	return products, err
}

func GetProductGroupByID(id uint) (*ProductGroup, error) {
	productGroup := ProductGroup{}
	err := database.DBCon.Db.
		QueryRow("SELECT id, name FROM product_groups WHERE id = $1", id).
		Scan(&productGroup.ID, &productGroup.Name)
	return &productGroup, err
}
