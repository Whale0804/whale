package controllers

import (
	"encoding/json"
	"github.com/githinkcn/whale/common"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/service"
	"github.com/githinkcn/whale/utils"
)

// Operations about Login
type LoginController struct {
	BaseController
}

// @Title register
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /register [post]
func (this *LoginController) Register() {
	loginRegisterDto := &entity.LoginRegisterDto{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginRegisterDto)

	userService := service.UserService{}
	if _, err := userService.FindByPhone(loginRegisterDto.Loginname); err == nil {
		this.Fail(common.ErrDupRecord, "用户已存在")
		return
	}

	loginService := service.LoginService{}
	id, _ := loginService.Register(loginRegisterDto)

	this.Resp(0, "success", map[string]interface{}{
		"id": id,
	})
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (this *LoginController) Login() {
	loginDto := &entity.LoginDto{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginDto)
	userService := service.UserService{}
	user, _ := userService.FindByPhone(loginDto.Loginname)
	if user.Id == 0 {
		this.Fail(common.ErrNoUser, "用户不存在")
		return
	}

	oldPassword, err := utils.Base64Decode(user.Password)

	if err != nil {
		this.Fail(common.ErrPass, "用户名或密码错误")
		return
	} else if oldPassword != loginDto.Password {
		this.Fail(common.ErrPass, "用户名或密码错误")
		return
	}

	if token, ok := this.GenToken(user.Id, user.Loginname); ok == nil {
		this.Resp(0, "success", map[string]interface{}{
			"token": token,
		})
	}

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [post]
func (this *LoginController) Logout() {

}
