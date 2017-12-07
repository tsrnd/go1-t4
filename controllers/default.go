package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Test() {
	this.InitFrontEndTemplate("frontend/user/login_register.tpl")
}

func (this *MainController) TestAdmin() {
	this.InitAdminTemplate("admin/product/create.tpl")
}

func (this *MainController) InitFrontEndTemplate(TplName string, ) {
	this.Layout = "frontend/master.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/layouts/header.tpl"
	this.LayoutSections["Footer"] = "frontend/layouts/footer.tpl"
	this.LayoutSections["Modal"] = "frontend/layouts/modal.tpl"
}

func (this *MainController) InitAdminTemplate(TplName string, ) {
	this.Layout = "admin/master.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "admin/layouts/header.tpl"
	this.LayoutSections["Footer"] = "admin/layouts/footer.tpl"
	this.LayoutSections["Aside"] = "admin/layouts/aside.tpl"
}
