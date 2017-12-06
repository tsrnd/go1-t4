package handlers

import (
	"github.com/goweb4/utils"
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
  userRequest, err := models.GetUserByUserName(GetAuthName(r)); if err != nil {
    fmt.Fprintln(w, err);
    return
  }
  order.UserID = userRequest.ID
  order.TotalMoney, err = strconv.ParseFloat(r.FormValue("total_money"), 64); if err != nil {
    fmt.Fprintln(w, err)
    return
  }
  DoStuff(r, &order)
  _, err = models.CreateOrder(order); if err != nil {
    fmt.Fprintln(w, err)
    return
  }
  fmt.Fprintln(w, "Order succeed")
}

/**
  * Add Form request to Order Products
  */
func DoStuff(r *http.Request, order *models.Order) (err error){
  for index, value := range r.PostForm["product_id"] {
    orderProduct := models.OrderProduct{}
    orderProduct.ProductID, err = utils.ConvertStrToUint(value); if err != nil {
      break
    }
    orderProduct.Quantity, err = utils.ConvertStrToUint(r.PostForm["quantity"][index]); if err != nil {
      break
    }
    order.OrderProducts = append(order.OrderProducts, orderProduct);
  }
  return err 
}

/**qqqq
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

/**
  * User update order's infor
  */
  func Checkout(w http.ResponseWriter, r *http.Request) {
    if GetAuthName(r) == "" {
      http.Redirect(w, r, "/login", 302)      
    }
    utils.GenerateTemplate(w, NewHomePageVars(r), "checkout")
    // fmt.Fprintln(w, "Need to be implement");
  }
