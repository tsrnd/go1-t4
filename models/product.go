package models

import (
	"github.com/goweb4/database"
)

type Product struct {
	Model
	Size          string         			`schema:"size"`
	Color         string         			`schema:"color"`
	Price         float64        			`schema:"price"`
	Name          string         			`schema:"name"`
	InStock       uint           			`schema:"in_stock"`
	GroupID       uint           			`schema:"group_id"`
	ProductGroup  *ProductGroup   		//belong To Product Group
	OrderProducts []OrderProduct 			//has many order products
	Images        []Image        			//has many image
}

func (product *Product) GetSchema() ([]interface{}) {
	return []interface{} {
		&product.ID,
		&product.Size,
		&product.Color,
		&product.Price,
		&product.InStock,
		&product.GroupID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
		&product.Name,
	}
}

func (product *Product) TableName() (string) {
	return "products"
}

func (product *Product) Test() {
	database.DBCon.Where("id = ?", 11).Find(product)
}

// func GetProducts() (products []Product) {
// 	err := database.DBCon.Find(&products).Error
// 	if err != nil {
// 		return products
// 	}
// 	return products
// }

// func GetProduct(id uint) (product Product, err error) {
// 	err = database.DBCon.Where("id = ?", id).Find(&product).Error
// 	database.DBCon.Model(&product).Association("Images").Find(&product.Images)

// 	return product, err
// }

// func UpdateProduct(product *Product) (errUpdate error) {
// 	errUpdate = database.DBCon.Save(&product).Error
// 	return errUpdate
// }

// func DeleteProduct(id uint) (errGet error) {
// 	product := Product{}
// 	product, errGet = GetProduct(id)
// 	if errGet == nil {
// 		errGet = database.DBCon.Delete(&product).Error
// 	}
// 	return errGet
// }

// func CreateProduct(product *Product) (proId uint, err error) {
// 	err = database.DBCon.Create(&product).Error
// 	if proId = 0; err == nil {
// 		proId = product.ID
// 	}
// 	return proId, err
// }
// func GetTrendProducts() (listProduct []Product) {

// 	rows, err := database.DBCon.Table("order_products").
// 		Select("product_id, sum(quantity) as total").
// 		Group("product_id").
// 		Order("total desc").
// 		Limit(4).
// 		Rows()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for rows.Next() {
// 		var id, quantity uint
// 		product := Product{}
// 		if err := rows.Scan(&id, &quantity); err != nil {
// 			fmt.Println(err)
// 		}
// 		database.DBCon.Where("id = ?", id).First(&product)
// 		listProduct = append(listProduct, product)
// 	}
// 	return listProduct
// }
