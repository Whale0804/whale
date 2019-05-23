// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/githinkcn/whale/controllers"
	"github.com/githinkcn/whale/utils"
	"strings"
)

func init() {
	var FilterAuth = func(ctx *context.Context) {
		authorization := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))
		if authorization == "" {
			ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
		}
		tokenString := strings.TrimSpace(authorization[len("Bearer "):])
		if _, isValid, err := utils.ParaseToken(tokenString); err == nil && !isValid {
			ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
		}
	}
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSBefore(FilterAuth),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
