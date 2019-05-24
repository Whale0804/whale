package dao

import (
	"github.com/githinkcn/whale/config"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/utils"
)

type UserDao struct{}

func (dao *UserDao) FindByUserName(name string) (user models.User, err error) {
	o := GetOrmer()
	err = o.Raw("select * from "+config.USER_TABLE+" where login_name = ?", name).QueryRow(&user)
	return user, err
}

func (dao *UserDao) FindByPhone(loginname string) (user models.User, err error) {
	o := GetOrmer()
	err = o.Raw("select * from "+config.USER_TABLE+" where login_name = ?", loginname).QueryRow(&user)
	return user, err
}

func (dao *UserDao) NewUser(dto *entity.UserAddDto) (id int, err error) {
	o := GetOrmer()
	user := &models.User{
		Id:        utils.GetId(),
		Loginname: dto.Loginname,
		Password:  utils.Base64Encode(dto.Password),
		Phone:     dto.Phone,
		Avatar:    dto.Avatar,
		DeptId:    dto.DeptId,
		Status:    dto.Status,
		Email:     dto.Email,
		Sex:       dto.Sex,
	}
	_, err = o.Insert(user)
	if err != nil {
		return 0, err
	}
	return int(user.Id), nil
}
