package test

import (
	"log"
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
	product.ProductGroup = productGroup
	CreateProductGroupTable()
	CreateProductTable()
	_, err := models.CreateProduct(product)
	if err != nil {
		t.Errorf("Can not create product")
	}

	DropProduct()
	DropProductGroup()
}

func Setup() {
	database.DBCon.Db, _ = database.DBConnection()
}

func TearDown() {
	database.DBCon.Db.Close()
}

func CreateProductTable() {
	_, err := database.DBCon.Db.Exec("CREATE TABLE products ( id serial, color varchar(32), size varchar(32), price varchar(32), name varchar(32), in_stock integer, group_id integer, created_at time without time zone, updated_at time without time zone, deleted_at time without time zone )")
	if err != nil {
		panic(err)
	}
}

func CreateProductGroupTable() {
	_, err := database.DBCon.Db.Exec("CREATE TABLE product_groups ( id serial, name varchar(32), created_at time without time zone, updated_at time without time zone, deleted_at time without time zone)")
	if err != nil {
		panic(err)
	}
}

func DropProduct() {
	if _, err := database.DBCon.Db.Exec("DROP TABLE products"); err != nil {
		log.Print(err)
	}
}

func DropProductGroup() {
	if _, err := database.DBCon.Db.Exec("DROP TABLE product_groups"); err != nil {
		log.Print(err)
	}
}
