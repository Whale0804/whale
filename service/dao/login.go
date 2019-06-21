package dao

import (
	"github.com/astaxie/beego"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/utils"
	"os"
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
	if err != nil {
		return 0, err
	}
	//用户文件夹关联
	uf := &models.UserFolder{
		Id:       utils.GetId(),
		UserId:   user.Id,
		FolderId: folderId,
	}
	_, err = o.Insert(uf)
	if err != nil {
		return 0, err
	}
	//用户问价文件夹文件关联表
	uff := &models.UserFolderFile{
		Id:       utils.GetId(),
		UserId:   user.Id,
		FolderId: folderId,
		Type:     0,
	}
	_, err = o.Insert(uff)
	if err != nil {
		return 0, nil
	}
	filePath := beego.AppConfig.String("whale_path") + folder.FolderName + "/"
	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return 0, nil
	}
	return int(user.Id), nil
}
