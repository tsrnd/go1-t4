package test

import (
	"testing"

	"github.com/goweb4/database"
	"github.com/goweb4/models"
)

func TestCreateProduct(t *testing.T) {
	Setup()
	defer TearDown()
	product := new(models.Product)
	product.Name = "test"
	product.Price = 10000
	productGroup := new(models.ProductGroup)
	productGroup.Name = "group1"
	product.ProductGroup = *productGroup
	database.DBCon.CreateTable(&models.ProductGroup{})
	database.DBCon.CreateTable(&models.Product{})
	_, err := models.CreateProduct(product)
	if err != nil {
		t.Errorf("Can not create product")
	}
	database.DBCon.DropTable(&productGroup)
	database.DBCon.DropTable(&product)
}

func Setup() {
	database.DBCon, _ = database.DBConnection()
}

func TearDown() {
	database.DBCon.Close()
}
