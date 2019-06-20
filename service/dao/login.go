package dao

import (
	"fmt"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/utils"
	"strconv"
)

type LoginDao struct{}

func (dao *LoginDao) Register(dto *entity.LoginRegisterDto) (id int, err error) {
	o := GetOrmer()
	//创建用户
	user := &models.User{
		Id:        utils.GetId(),
		Loginname: dto.Loginname,
		Password:  utils.Base64Encode(dto.Password),
		Phone:     dto.Loginname,
	}
	//插入用户信息
	_, err = o.Insert(user)
	if err != nil {
		return 0, err
	}
	//文件夹信息
	folderId := utils.GetId()
	folder := &models.Folder{
		Id:         folderId,
		FolderName: strconv.Itoa(user.Id),
		ParentId:   0,
	}
	//插入文件夹信息
	_, err = o.Insert(folder)
	//用户文件夹关联
	uf := &models.UserFolder{
		Id:       utils.GetId(),
		UserId:   user.Id,
		FolderId: folderId,
	}
	_, err = o.Insert(uf)
	//用户问价文件夹文件关联表
	uff := &models.UserFolderFile{
		Id:       utils.GetId(),
		UserId:   user.Id,
		FolderId: folderId,
		Type:     0,
	}
	_, err = o.Insert(uff)
	if err != nil {
		fmt.Println(err)
	}
	return int(user.Id), nil
}
