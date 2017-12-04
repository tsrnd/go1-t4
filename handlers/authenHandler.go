package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goweb4/models"
	"github.com/goweb4/utils"
	"github.com/stretchr/objx"
)

type HomePageVars struct {
	Name         string
	Message      string
	RegisterInfo RegisterInfo
	PageTitle    string
	ProductGroup []models.ProductGroup
	Product 		 models.Product
}

/**
 * HomePageVars constructor
 */
func NewHomePageVars(r *http.Request) HomePageVars {
	var homePageVars HomePageVars
	homePageVars.ProductGroup = GetProductGroups()
	homePageVars.Name = GetAuthName(r)

	return homePageVars
}

func Index(w http.ResponseWriter, r *http.Request) {
	HomeVars := NewHomePageVars(r)
	HomeVars.PageTitle = "Home Page"
	utils.GenerateTemplate(w, HomeVars, "index")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err != nil {
		HomeVars := NewHomePageVars(r)
		HomeVars.Message = utils.ShowMessage(w, r, "login")
		HomeVars.PageTitle = "Login"
		utils.GenerateTemplate(w, HomeVars, "login_register", "login", "register")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var info models.User
	err := utils.MapFormValues(&info, r)
	if err != nil {
		fmt.Println("cannot decode login info: ", err)
	} else {
		if utils.CheckCredential(info) {
			setSession(info.UserName, w)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
		} else {
			mess := "Sorry, this does not match our records. Check your spelling and try again."
			utils.SetMessage(w, mess, "login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}

func setSession(userName string, response http.ResponseWriter) {
	authCookieValue := objx.New(map[string]interface{}{
		"name": userName,
	}).MustBase64()
	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name:    "auth",
		Expires: expiration,
		Value:   authCookieValue})
}

func clearSession(response http.ResponseWriter) {
	http.SetCookie(response, &http.Cookie{
		Name:    "auth",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	})
}

/**
 * Get product_group for header
 */
func GetProductGroups() []models.ProductGroup {
	productGroups, err := models.GetProductGroups()
	if err != nil {
		fmt.Println("err")
	}
	return productGroups
}

func GetAuthName(r *http.Request) string {
	name := ""
	if authCookie, err := r.Cookie("auth"); err == nil {
		var cookieData interface{}
		cookieData = objx.MustFromBase64(authCookie.Value)
		name = cookieData.(objx.Map)["name"].(string)
	}
	return name
}
