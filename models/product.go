package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Size    string  `schema:"size"`
	Color   string  `schema:"color"`
	Price   float64 `schema:"price"`
	Name    string  `schema:"name"`
	InStock uint    `schema:"in_stock"`
	GroupID uint    `schema:"group_id"`
}

func GetProducts() (products []Product, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Find(&products).Error
	})
	return products, err
}

func GetProduct(id uint) (product Product, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("id = ?", id).Find(&product).Error
	})
	return product, err
}

func UpdateProduct(product *Product) (errUpdate error) {
	WithConnectionDB(func(db *database.DB) {
		errUpdate = db.Save(&product).Error
	})
	return errUpdate
}

func DeleteProduct(id uint) (errGet error) {
	product := Product{}
	WithConnectionDB(func(db *database.DB) {
		product, errGet = GetProduct(id)
		if errGet == nil {
			errGet = db.Delete(&product).Error
		}
	})
	return errGet
}

func CreateProduct(product *Product) (proId uint, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Create(&product).Error
		if proId = 0; err == nil {
			proId = product.ID
		}
	})
	return proId, err
}
