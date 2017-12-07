package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/goweb4/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/goweb4/controllers:UserController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/users/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
