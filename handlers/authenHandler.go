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
	Name          string
	Message       string
	RegisterInfo  RegisterInfo
	PageTitle     string
	ProductGroup  []models.ProductGroup
	Products      []models.Product
	LatesProducts []models.Product
	BestSeller    []models.Product
	Paginator     utils.Paginator
	User          models.User
	Product       models.Product
	Payments      []models.Payment
}

type Handler struct {
	Variable Authenticate
}

type Authenticate interface {
	NewHomePageVars(*http.Request) HomePageVars
	ShowMessage(w http.ResponseWriter, r *http.Request) string
	GenerateTemplate(w http.ResponseWriter, h HomePageVars)
}

type Vars struct{}

func (v *Vars) NewHomePageVars(r *http.Request) HomePageVars {
	var homePageVars HomePageVars
	homePageVars.ProductGroup, _ = models.GetProductGroups()
	homePageVars.Name = GetAuthName(r)

	return homePageVars
}

func (v *Vars) ShowMessage(w http.ResponseWriter, r *http.Request) string {
	return utils.ShowMessage(w, r, "login")
}

func (v *Vars) GenerateTemplate(w http.ResponseWriter, h HomePageVars) {
	utils.GenerateTemplate(w, h, "login_register", "login", "register")
}

func Index(w http.ResponseWriter, r *http.Request) {

	// // HomeVars := NewHomePageVars(r)
	// HomeVars.LatesProducts = models.GetLatestProduct(4)
	// HomeVars.BestSeller = models.GetTrendProducts(4)
	// HomeVars.PageTitle = "Home Page"
	// utils.GenerateTemplate(w, HomeVars, "index")
}

func IndexAdmin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/product/add", http.StatusSeeOther)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	HomeVars := h.Variable.NewHomePageVars(r)
	if _, err := r.Cookie("auth"); err != nil {
		HomeVars.Message = h.Variable.ShowMessage(w, r)
		HomeVars.PageTitle = "Login"
		h.Variable.GenerateTemplate(w, HomeVars)
	} else if utils.IsAdminRole(HomeVars.Name) {
		http.Redirect(w, r, "/adminIndex", 302)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var info models.User
	err := utils.MapFormValues(&info, r)
	if err != nil {
		fmt.Println("cannot decode login info: ", err)
		return
	}
	if utils.CheckCredential(info) {
		if utils.IsAdminRole(info.UserName) {
			setSession(info.UserName, w)
			http.Redirect(w, r, "/adminIndex", http.StatusSeeOther)
		} else {
			setSession(info.UserName, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		mess := "Sorry, this does not match our records. Check your spelling and try again."
		utils.SetMessage(w, mess, "login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
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

func GetAuthName(r *http.Request) string {
	name := ""
	if authCookie, err := r.Cookie("auth"); err == nil {
		var cookieData interface{}
		cookieData = objx.MustFromBase64(authCookie.Value)
		name = cookieData.(objx.Map)["name"].(string)
	}
	return name
}

func ContactUs(w http.ResponseWriter, r *http.Request) {
	// utils.GenerateTemplate(w, NewHomePageVars(r), "contact")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	// utils.GenerateTemplate(w, NewHomePageVars(r), "about")
}
