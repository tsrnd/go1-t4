package controllers

import (
	"github.com/goweb4/models"
	"net/http"
	"github.com/astaxie/beego"
	"github.com/goweb4/lib"
)

type AuthenController struct {
	ExtendController
}

// URLMapping ...
func (c *AuthenController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("LoginHandler", c.LoginHandler)
}

// Login ...
// @router /login [get]
func (c *AuthenController) Login() {
	if (c.IsLogin()) {
		c.Ctx.Redirect(302, c.URLFor("UserController.Show"))
		return
	}
	InitFrontEndTemplate(&c.Controller, "frontend/user/login.tpl")
}

// LoginHandler ...
// @router /login [post]
func (c *AuthenController) LoginHandler() {
	username := c.GetString("username")
	password := c.GetString("password")
	flash := beego.NewFlash()
	user, err := lib.CheckCredentials(username, password)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		InitFrontEndTemplate(&c.Controller, "frontend/user/login.tpl")
		return
	}
	flash.Success("Login succeed")
	flash.Store(&c.Controller)
	c.SetLogin(user)
	c.Redirect(c.URLFor("UserController.Show", ":id", user.Id), http.StatusSeeOther)
	c.Redirect("/users", http.StatusSeeOther)
}

func (c *AuthenController) IsLogin() (bool) {
	if c.Session.Get("uid") != nil && c.Session.Get("uid").(int64) > 0 {
		return true
	}
	return false
}

func (c *AuthenController) SetLogin(user *models.User) {
	c.Session.Set("uid", user.Id)
}
