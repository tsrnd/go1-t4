package controllers

import (
	"github.com/astaxie/beego"
)

//  ProductGroupsController operations for ProductGroups
type ProductGroupsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductGroupsController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
func (c *ProductGroupsController) GetAll() {
}