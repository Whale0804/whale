package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/githinkcn/whale/utils"
	"strings"
)

type BaseController struct {
	beego.Controller
}

//封装返回体
func (self *BaseController) Resp(code int, msg string, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	if len(data) >= 1 {
		out["data"] = data[0]
	}
	if len(data) >= 2 {
		out["total"] = data[1]
	}
	self.Data["json"] = out
	self.ServeJSON()
}

//生成token
func (this *BaseController) GenToken(Uid int, Uname string) (string, error) {
	return utils.GenToken(Uid, Uname)
}

//验证token
func (this *BaseController) ValidToken() (*utils.WhaleClaims, bool, error) {
	authorization := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	if authorization == "" {
		return nil, false, errors.New("Authorization is empty")
	}
	tokenString := strings.TrimSpace(authorization[len("Bearer "):])
	if claims, isValid, err := utils.ParaseToken(tokenString); err == nil && isValid {
		return claims, true, nil
	}
	return nil, false, errors.New("Authorization invalid")
}

//
