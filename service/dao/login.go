package dao

import (
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/utils"
)

type LoginDao struct{}

func (dao *LoginDao) Register(dto *entity.LoginRegisterDto) (id int, err error) {
	o := GetOrmer()
	user := &models.User{
		Id:        utils.GetId(),
		Loginname: dto.Loginname,
		Password:  utils.Base64Encode(dto.Password),
		Phone:     dto.Loginname,
	}
	_, err = o.Insert(user)
	if err != nil {
		return 0, err
	}
	return int(user.Id), nil
}
