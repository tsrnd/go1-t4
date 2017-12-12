package models

import (
	"fmt"
	"github.com/goweb4/database"	
	"errors"
)

const PENDING_STATUS = "Pending"
const APPROVED_STATUS = "Approved"
const CANCELED_STATUS = "Canceled"

type Order struct {
	Model
	UserID        uint           `schema:"user_id"`
	TotalMoney    float64        `schema:"total_money"`
	Status        string         `schema:"status"`
	PaymentID     uint           `schema:"payment_id"`
	Payment       *Payment       //has one payment method
	OrderProducts []OrderProduct //has many order products
	User          *User          //belong to user
}

// CreateOrder using for create Order and checkou Product
func CreateOrder(order Order) (orderID uint, err error) {
	// begin transaction 
	trans, err := database.DBCon.Db.Begin()
	if err != nil {
		// return if error, can not begin transaction
		fmt.Println(err)
		return
	}
	// Insert Order into table and get Id Order
	err = trans.QueryRow(`INSERT INTO orders (user_id, total_money, status) 
	VALUES ($1,$2,$3) RETURNING id;`, order.UserID, order.TotalMoney, order.Status).Scan(&order.ID)

	if err != nil {
		// rollback database and return when can't insert Order
		trans.Rollback()
		fmt.Println("Error when insert order! rollbacked!")
		return
	}

	// Prepare insert order_products table
	ins, err := trans.Prepare(`INSERT INTO order_products (order_id, product_id, quatity) 
	VALUES ($1, $2, $3);`)
	if err != nil {
		// return errors when cant prepare
		fmt.Println(err)
		return
	}
	defer ins.Close()
	// Excute insert order_products
	for _,v := range order.OrderProducts {
		v.OrderID = order.ID
		if _, err = ins.Exec(v.OrderID, v.ProductID, v.Quantity); err != nil {
			trans.Rollback() // return and rollback if errors when insert
			fmt.Println(err)
			return 
		} else {
			err = CheckoutAfterCreate(&v); if err != nil {
				fmt.Println(err)
				trans.Rollback()// return and rollback if errors when checkout product table
				return
			}
		}
	}
	// Finish transaction
	err = trans.Commit()
	return orderID, err
}

func CheckoutAfterCreate(orderProduct *OrderProduct) (err error) {
	var product Product
	// get field 'in stock' in product 
	query := "SELECT id, in_stock FROM products WHERE id = $1"
	err = database.DBCon.Db.QueryRow(query, orderProduct.ProductID).Scan(&product.ID, &product.InStock)
	if int(orderProduct.Quantity) > int(product.InStock) {
		// return if not enough 
		err = errors.New("Instock is not enough")
		return err
	}
	// Update field 'in stock' of product
	query = `UPDATE products SET in_stock=$1 WHERE id = $2`
	_, err = database.DBCon.Db.Exec(query, product.InStock - orderProduct.Quantity, product.ID)  
	if err != nil {  
		// return if update fail
		fmt.Println(err)
		return err
	}

	return err
}

