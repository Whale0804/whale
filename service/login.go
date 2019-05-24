package service

import (
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/service/dao"
)

type LoginService struct {
	dao *dao.LoginDao
}

func (this *LoginService) Register(dto *entity.LoginRegisterDto) (id int, err error) {
	return this.dao.Register(dto)
}
