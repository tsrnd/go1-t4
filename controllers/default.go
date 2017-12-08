package controllers

import (
	"github.com/goweb4/models"
	"html/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type ExtendController struct {
	beego.Controller
	Session session.Store
}

func (c *ExtendController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *ExtendController) Test() {
	InitFrontEndTemplate(this, "frontend/user/login_register.tpl")
}

func (this *ExtendController) TestAdmin() {
	InitAdminTemplate(this, "admin/product/create.tpl")
}

func InitFrontEndTemplate(this *ExtendController, TplName string) {
	this.Layout = "frontend/master.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/layouts/header.tpl"
	this.LayoutSections["Footer"] = "frontend/layouts/footer.tpl"
	this.LayoutSections["Modal"] = "frontend/layouts/modal.tpl"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	if this.IsLogin() {
		userInfor := models.User{}
		data, _ := models.GetUserById(this.Session.Get("uid").(int64))
		err := userInfor.HideColumns(data); if err != nil {
			return
		}
		this.Data["User"] = userInfor
	}
}

func InitAdminTemplate(this *ExtendController, TplName string) {
	this.Layout = "admin/master.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "admin/layouts/header.tpl"
	this.LayoutSections["Footer"] = "admin/layouts/footer.tpl"
	this.LayoutSections["Aside"] = "admin/layouts/aside.tpl"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
}

func (this *ExtendController) Prepare() {
	this.Session = this.StartSession()
}

func (this *ExtendController) IsLogin() (bool) {
	if this.Session.Get("uid") != nil && this.Session.Get("uid").(int64) > 0 {
		return true
	}
	return false
}
