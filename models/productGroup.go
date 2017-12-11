package models

type ProductGroup struct {
	Model
	Name     string    `schema:"name"`
	Products []Product //has many products
}

func GetProductGroups() (productGroups []ProductGroup, err error) {
	// err = database.DBCon.Find(&productGroups).Error
	return productGroups, err
}

// func GetProductsByGroupID(id uint) (products []Product, err error) {
// 	err = database.DBCon.Where("group_id = ?", id).Find(&products).Error
// 	for i, _ := range products {
// 		database.DBCon.Model(products[i]).Related(&products[i].Images)
// 	}
// 	return products, err
// }
