package controllers

import (
	
)

//  ProductGroupsController operations for ProductGroups
type ProductGroupsController struct {
	ExtendController
}

// URLMapping ...
func (c *ProductGroupsController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
func (c *ProductGroupsController) GetAll() {
}
