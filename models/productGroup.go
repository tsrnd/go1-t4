package models

import (
	"fmt"
	"github.com/goweb4/database"
	"reflect"
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
	return productGroups
}

// Get products by group
func GetProductsByGroupID(id uint) (products []Product, err error) {
	prList := make(map[int]*Product)
	query := `SELECT id, name, size, price, in_stock, color FROM products WHERE group_id = $1`
	rows, err := database.DBCon.Db.Query(query, id)
		for rows.Next() {
			product := &Product{}
			err = rows.Scan(&product.ID, &product.Name, &product.Size, &product.Price, &product.InStock, &product.Color)
			if err != nil {
				return
			}
			prList[int(product.ID)] = product
		}
	defer rows.Close()

	rows, err = database.DBCon.Db.Query(`SELECT url, product_id FROM images`)
	if err != nil {
		return
	}
	for rows.Next() {
		i := &Image{}
		rows.Scan(&i.URL, &i.ProductId)
		// fmt.Println(prList[int(i.ProductId)])
		prList[int(i.ProductId)].Images = append(prList[int(i.ProductId)].Images, i)
	}
 	defer rows.Close()
	
	var pro []Product
	for _, p := range prList {
		fmt.Println(reflect.TypeOf(p))

		products = append(products, p)
	}

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
		image := &Image{}
		err = rows.Scan(&image.URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		product.Images = append(product.Images, image)
	}
	rows.Close()
}
