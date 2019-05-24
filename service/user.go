package service

import (
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/service/dao"
)

type UserService struct {
	dao *dao.UserDao
}

func (us *UserService) FindByUserName(name string) (user models.User, err error) {
	return us.dao.FindByUserName(name)
}

func (us *UserService) NewUser(dto *entity.UserAddDto) (id int, err error) {
	return us.dao.NewUser(dto)
}
