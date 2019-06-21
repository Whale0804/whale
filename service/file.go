package service

import (
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
	"github.com/githinkcn/whale/service/dao"
)

type FileService struct {
	dao *dao.FileDao
}

func (this *FileService) FindTopFolderByUserId(userId int) (folder models.Folder, err error) {
	return this.dao.FindTopFolderByUserId(userId)
}

func (this *FileService) SaveFile(dto *entity.FinishUploadDto) (file models.File, err error) {
	return this.dao.SaveFile(dto)
}
