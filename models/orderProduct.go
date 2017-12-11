package models

type OrderProduct struct {
	Model
	OrderID   uint    `schema:"order_id"`
	ProductID uint    `schema:"product_id"`
	Quantity  uint    `schema:"quantity"`
	Product   Product //belong to product
	Order     Order   //belong to order
}

// func CreateOrderProduct(orderProduct *OrderProduct) (orderProID uint, err error){
// 	err = database.Tx.Create(&orderProduct).Error
// 	if orderProID == 0 && err == nil {
// 		orderProID = orderProduct.ID
// 	}
// 	return orderProID, err
// }

// func (op *OrderProduct) AfterCreate(tx *gorm.DB) (err error) {
// 	product := new(Product)
// 	database.DBCon.Where("id = ?", op.ProductID).First(&product)
// 	if op.Quantity > product.InStock {
// 		err = errors.New("This product is not enough quantity in stock")
// 	} else {
// 		product.InStock -= op.Quantity
// 		tx.Save(&product)
// 	}
// 	return err
// }
