package handlers

import (
	"github.com/goweb4/utils"
	"github.com/goweb4/database"
  "fmt"
  "net/http"
  "github.com/goweb4/models"
  "github.com/gorilla/mux"
  "strconv"
)

/**
  * User create new order 
  */
func CreateOrder(w http.ResponseWriter, r *http.Request) {
  
}

/**
  * User create new order 
  */
func StoreOrder(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  order := models.Order{}
  userRequest, err := models.GetUserByUserName(r.FormValue("username")); if err != nil {
    fmt.Fprintln(w, err);
    return
  }
  order.UserID = userRequest.ID
  order.TotalMoney, err = strconv.ParseFloat(r.FormValue("total_money"), 64); if err != nil {
    fmt.Fprintln(w, err)
    return
  }
  //begin transaction
  database.Tx = database.DBCon.Begin()
  orderID, err := models.CreateOrder(order); if err != nil {
    fmt.Println(w, err)
    return
  }
  status := true
  for index, value := range r.PostForm["product_id"] {
    orderProduct := new(models.OrderProduct)
    orderProduct.OrderID = orderID
    orderProduct.ProductID, err = utils.ConvertStrToUint(value); if err != nil {
      fmt.Fprintln(w, err)
      status = false
      break
    }
    orderProduct.Quantity, err = utils.ConvertStrToUint(r.PostForm["quantity"][index]); if err != nil {
      fmt.Fprintln(w, err)
      status = false
      break
    }
    _, err = models.CreateOrderProduct(orderProduct); if err != nil {
      fmt.Fprintln(w, err)
      status = false
      break
    }
  }
  if !status {
    database.Tx.Rollback()
    fmt.Fprintln(w, err);
    return
  }
  database.Tx.Commit()
  fmt.Fprintln(w, "Order succeed")
}

/**
  * Show form order's edit 
  */
func EditOrder(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  orderId, _ := strconv.Atoi(vars["id"]);
  order, err := models.GetOrder(orderId); if err != nil {
    fmt.Fprintln(w, err);
  }
  
	fmt.Fprintln(w, order);
}

/**
  * User update order's infor
  */
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}
