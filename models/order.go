package models

const PENDING_STATUS = "Pending"
const APPROVED_STATUS = "Approved"
const CANCELED_STATUS = "Canceled"

type Order struct {
	UserID        uint           `schema:"user_id"`
	TotalMoney    float64        `schema:"total_money"`
	Status        string         `schema:"status"`
	PaymentID     uint           `schema:"payment_id"`
	Payment       Payment        //has one payment method
	OrderProducts []OrderProduct //has many order products
	User          User           //belong to user
}

func (order *Order) GetRelationship() map[string]interface{} {
	relationship := map[string]interface{}{
		"OrderProducts": &order.OrderProducts,
		"User":          &order.User,
		"Payment":       &order.Payment,
	}
	return relationship
}

// func GetOrder(id int) (order Order, err error) {
// 	err = database.DBCon.Where("id = ?", id).Find(&order).Error
// 	return order, err
// }

// func GetOrdersByUser(id int) (orders []Order, err error) {
// 	err = database.DBCon.Where("user_id = ?", id).Find(&orders).Error
// 	return orders, err
// }

// func CreateOrder(order Order) (orderID uint, err error) {
// 	err = database.DBCon.Create(&order).Error
// 	if orderID == 0 && err == nil {
// 		orderID = order.ID
// 	}

// 	return orderID, err
// }
