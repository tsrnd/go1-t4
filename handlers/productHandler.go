package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/goweb4/models"
	"github.com/goweb4/utils"
)

/**
 * Index Product
 */
func IndexProduct(w http.ResponseWriter) {
	fmt.Fprintln(w, "Need to be implement")
}

/**
 * Show product
 */
func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	product, err := models.GetProduct(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, product)
	// utils.GenerateTemplate(w, handlers.HomeVars)
}

/**
 * Show product by group
 */
func ShowProductGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	query := r.URL.Query()
	page, _ := strconv.ParseInt(query.Get("page"), 10, 32)

	HomeVars := NewHomePageVars(r)

	products, err := models.GetProductsByGroupID(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
	}

	paginator := utils.Paginate(len(products), 12, int(page))

	HomeVars.Products = products[paginator.Start:paginator.End]
	HomeVars.Paginator = paginator

	utils.GenerateTemplate(w, HomeVars, "product")
}

/**
 * Show form create new product
 */
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	productGroups, errGet := models.GetProductGroups()
	if errGet != nil {
		http.Error(w, errGet.Error(), http.StatusInternalServerError)
		return
	}
	utils.GenerateTemplateAdmin(w, productGroups, "add_product")
}

/**
 * Admin create new product
 */
func StoreProduct(w http.ResponseWriter, r *http.Request) {
	product := new(models.Product)
	errMap := utils.MapFormValues(product, r)
	if errMap != nil {
		fmt.Println("Error occur when mapping value: ", errMap)
		fmt.Fprintln(w, errMap)
		return
	}

	id, errCreate := models.CreateProduct(product)
	if errCreate != nil {
		fmt.Fprintln(w, errCreate)
		return
	}

	//get image from form file request
	file, header, err := r.FormFile("image")
	if err == nil {
		fileName, _, err := utils.HandleImage(file, header, id, models.IMG_BASE_URL)
		if err != nil {
			http.Redirect(w, r, "/product/add", http.StatusBadRequest)
			return
		}
		image := models.Image{}
		image.Name = fileName
		image.ProductId = id
		image.URL = strings.Join([]string{models.IMG_BASE_URL, fileName + ".jpg"}, "/")
		err = models.StoreImage(&image)
		if err != nil {
			http.Redirect(w, r, "/product/add", http.StatusBadRequest)
			return
		}
	}

	http.Redirect(w, r, "/product/"+fmt.Sprint(id), http.StatusFound)
}

/**
 * Show form product's edit
 */
func EditProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	product, err := models.GetProduct(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, product)
}

/**
 * User update product's infor
 */
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	product, err := models.GetProduct(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	errMap := utils.MapFormValues(&product, r)
	if errMap != nil {
		log.Fatal(errMap)
		fmt.Fprintln(w, errMap)
		return
	}
	errUpdate := models.UpdateProduct(&product)
	if errUpdate != nil {
		fmt.Fprintln(w, errUpdate)
		return
	}
	fmt.Fprintln(w, "Update this product success")
}

/**
 * Delete product
 */
func DestroyProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	err := models.DeleteProduct(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	} else {
		fmt.Fprintln(w, "Delete this product success")
	}
}

 /**
  * Detail product
  */
  func DetailProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.ParseUint(vars["id"], 10, 32);
    product, err := models.GetProduct(uint(id)); if err != nil {
      fmt.Fprintln(w, err);
    }

    products, err := models.GetProductsByGroupID(product.GroupID); if err != nil {
      fmt.Fprintln(w, err);
    }
    fmt.Println(product)
    fmt.Println(products)

    HomeVars := NewHomePageVars(r)
    HomeVars.Product = product
    HomeVars.Products = products[:3]
    utils.GenerateTemplate(w, HomeVars, "product_detail")
  }
