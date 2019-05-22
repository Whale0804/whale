package controllers

import ()

type LoginController struct {
	BaseController
}

// @Title 用户登陆
// @Description 用户登陆 http://localhost:8080/api/v1/user/1/update
// @Param   username
// @Param   password
// @Success 2000
// @Failure 4001 User not found
// @router / [post]
func (this *LoginController) Login() {

}