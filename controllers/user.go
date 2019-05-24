package controllers

import (
	"encoding/json"
	"github.com/githinkcn/whale/common"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/service"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	userAddDto := &entity.UserAddDto{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userAddDto)
	userService := service.UserService{}
	if _, err := userService.FindByUserName(userAddDto.Username); err == nil {
		this.Fail(common.ErrDupRecord, "用户已存在")
		return
	}
	id, _ := userService.NewUser(userAddDto)

	this.Resp(0, "success", map[string]interface{}{
		"id": id,
	})
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {

}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {

}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {

}
