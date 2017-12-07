package controllers

import (
	"github.com/astaxie/beego"
)

//  UsersController operations for Users
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Show", c.Show)
}

// @router /users/:id [get]
func (c *UserController) Show() {
	
}
