package models

import (
	"github.com/goweb4/database"
)

type Product struct {
	Model
	Size          string         `schema:"size"`
	Color         string         `schema:"color"`
	Price         float64        `schema:"price"`
	Name          string         `schema:"name"`
	InStock       uint           `schema:"in_stock"`
	GroupID       uint           `schema:"group_id"`
	ProductGroup  *ProductGroup  //belong To Product Group
	OrderProducts []OrderProduct //has many order products
	Images        []Image        //has many image
}

func (product *Product) GetSchema() (map[string]interface{}) {
	return map[string]interface{} {
		"id": &product.ID,
		"size": &product.Size,
		"color": &product.Color,
		"price": &product.Price,
		"in_stock": &product.InStock,
		"group_id": &product.GroupID,
		"created_at": &product.CreatedAt,
		"updated_at": &product.UpdatedAt,
		"deleted_at": &product.DeletedAt,
		"name": &product.Name,
	}
}

func (product *Product) TableName() string {
	return "products"
}

// func GetProducts() (products []Product) {
// 	err := database.DBCon.Find(&products).Error
// 	if err != nil {
// 		return products
// 	}
// 	return relationship
// }

// func GetProducts() (products []Product) {
// err := database.DBCon.Find(&products).Error
// if err != nil {
// 	return products
// }
// return products
// }

func GetProduct(id uint) (product Product, err error) {
	err = database.DBCon.Where("id = ?", id).Find(&product)
	if err != nil {
		return
	}
	rows, err := database.DBCon.Db.
		Query("SELECT * FROM images WHERE product_id = $1", product.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		image := Image{Product: &product}
		err = rows.Scan(
			&image.ID, &image.Name, &image.URL, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt, &image.ProductId,
		)
		if err != nil {
			return
		}
		product.Images = append(product.Images, image)
	}
	return
}

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
func GetTrendProducts() (listProduct []Product) {

	// rows, err := database.DBCon.Table("order_products").
	// 	Select("product_id, sum(quantity) as total").
	// 	Group("product_id").
	// 	Order("total desc").
	// 	Limit(4).
	// 	Rows()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for rows.Next() {
	// 	var id, quantity uint
	// 	product := Product{}
	// 	if err := rows.Scan(&id, &quantity); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	database.DBCon.Where("id = ?", id).First(&product)
	// 	listProduct = append(listProduct, product)
	// }
	return listProduct
}

func GetLatestProduct(limit int) (products []Product) {
	rows, err := database.DBCon.Db.Query("SELECT id, size, name, price from products ORDER BY id DESC limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Size, &product.Name, &product.Price)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return products
}
