package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nfnt/resize"
	"os"
	"net/http"
	"fmt"
	"strings"
	"path/filepath"
	"log"
	"image"
	"image/jpeg"
	"github.com/goweb4/database"
)

type Image struct {
	gorm.Model
	Name 		string		`schema:"name"`
	URL  		string		`schema:"url"`
	ProductId	uint		`schema:"product_id"`
}

const IMG_BASE_URL = ""

func StoreImage(w http.ResponseWriter, r *http.Request, id uint) (error){
	imageResult := Image{}
	//connect to db
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
		return errConnection
	}
	defer db.Close()
	//get image from form file request
	file, header, err := r.FormFile("image")
	if err != nil {
		return err
	}

	fileName := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
	fmt.Println(fileName)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	m := resize.Resize(300, 300, img, resize.Lanczos3)
	
	out, err := os.Create(fileName + ".jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	// Encode into jpeg
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return err
	}
	imageResult.Name = fileName
	imageResult.ProductId = id
	imageResult.URL = IMG_BASE_URL
	
	newImage := db.Create(&imageResult)
	errCreateImage := newImage.Error

	return errCreateImage
}
