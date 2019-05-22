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

type TokenCheckController struct {
	Id       int
	Username string
	RawToken string
	BaseController
}

// 固定返回的json数据格式
// code: 错误码
// msg: 错误信息
// data: 返回数据
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
func (this *BaseController) GenToken(Uid int) (string, error) {
	return utils.GenToken(Uid)
}

//验证token
func (this *BaseController) ValidToken() (int, bool, error) {
	authorization := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	if authorization == "" {
		return 0, false, errors.New("Authorization is empty")
	}
	if claims, isValid, err := utils.ParaseToken(authorization); err == nil && isValid {
		return claims.Uid, true, nil
	}
	return 0, false, errors.New("Authorization invalid")
}
