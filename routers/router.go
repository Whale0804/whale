// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/githinkcn/whale/config"
	"github.com/githinkcn/whale/controllers"
	"github.com/githinkcn/whale/utils"
	"strings"
)

func init() {
	//https://blog.csdn.net/mirage003/article/details/87865582
	var FilterAuth = func(ctx *context.Context) {
		if flag, _ := beego.AppConfig.Bool("auth"); flag == true {
			authorization := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))
			if authorization == "" {
				ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
			}
			tokenString := strings.TrimSpace(authorization[len("Bearer "):])
			if userInfo, isValid, err := utils.ParaseToken(tokenString); err == nil && isValid {
				fmt.Println(userInfo.Uname)
				if config.Cache.IsExist("login:" + userInfo.Uname) {
					v := string(config.Cache.Get("login:" + userInfo.Uname).([]byte))
					if redisUserInfo, isValidRedis, errRedis := utils.ParaseToken(v); errRedis == nil && isValidRedis {
						if redisUserInfo.Uname != userInfo.Uname || redisUserInfo.Uid != userInfo.Uid {
							ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
						}
					} else {
						ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
					}
				} else {
					ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
				}

			} else {
				ctx.Output.JSON(map[string]interface{}{"code": 401, "msg": "请登录后访问"}, true, true)
			}
		}
	}
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(
				&controllers.FileController{},
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
