package controllers

import (
	"time"
)

// Operations about Login
type LoginController struct {
	BaseController
}

// @Title register
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /register [get]
func (this *LoginController) Register() {

}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (this *LoginController) Login() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")

	if username == "mzk" && password == "123456" {
		if token, ok := this.GenToken(int(time.Now().Unix()), username); ok == nil {
			this.Resp(0, "success", map[string]interface{}{
				"token": token,
			})
		}

	}

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (this *LoginController) Logout() {

}
