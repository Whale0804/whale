package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/githinkcn/whale/common"
	"github.com/githinkcn/whale/utils"
	"reflect"
	"strings"
)

type BaseController struct {
	ctx *context.Context
	beego.Controller
	username string
	userId   int
}

func (this *BaseController) Prepare() {

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

// 解析并验证表单，返回第一个错误信息
func (b *BaseController) ParseAndValidateFirstErr(obj interface{}) error {
	if err := b.ParseForm(obj); err != nil {
		return err
	}
	valid := &validation.Validation{}
	if v, _ := valid.Valid(obj); !v {
		// stuctTag
		tags := make(map[string]interface{})
		s := reflect.TypeOf(obj).Elem()
		for i := 0; i < s.NumField(); i++ {
			tags[s.Field(i).Name] = s.Field(i).Tag.Get("form")
		}
		for _, err := range valid.Errors {
			return errors.New(tags[err.Field].(string) + ":" + err.Message)
		}
	}

	return nil
}

func (self *BaseController) Fail(errs *common.ControllerError, moreErrInfo ...string) {

	self.Data["json"] = errs
	fmt.Println(self.Data)
	errs.Moreinfo = ""
	for _, v := range moreErrInfo {
		errs.Moreinfo += v
	}
	self.ServeJSON()
}
