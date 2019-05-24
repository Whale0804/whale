package dao

import (
	"github.com/githinkcn/whale/config"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/utils"
)

type UserDao struct {
}

func (dao *UserDao) FindByUserName(name string) (user models.User, err error) {
	o := GetOrmer()
	err = o.Raw("select * from "+config.USER_TABLE+" where user_name = ?", name).QueryRow(&user)
	return user, err
}

func (dao *UserDao) NewUser(dto *entity.UserAddDto) (id int, err error) {
	passeord := utils.Base64Encode(dto.Password)
	o := GetOrmer()
	users := &models.User{
		Id:        utils.GetId(),
		Loginname: dto.Loginname,
		Password:  passeord,
		Phone:     dto.Phone,
		Avatar:    dto.Avatar,
		DeptId:    dto.DeptId,
		Status:    dto.Status,
		Email:     dto.Email,
		Sex:       dto.Sex,
	}
	_, err = o.Insert(users)
	if err != nil {
		return 0, err
	}
	return int(users.Id), nil
}
