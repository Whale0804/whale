package dao

import (
	"github.com/githinkcn/whale/config"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/models"
)

type FileDao struct{}

func (dao *FileDao) FindTopFolderByUserId(userId int) (folder models.Folder, err error) {
	o := GetOrmer()
	err = o.Raw("SELECT * FROM "+config.FOLDER_TABLE+" f WHERE f.id in (SELECT uf.folder_id AS folder_id FROM "+config.USER_FOLDER_TABLE+" uf WHERE uf.user_id = ?) AND f.parent_id = '0'", userId).QueryRow(&folder)
	return folder, err
}

func (dao *FileDao) SaveFile(dto *entity.FinishUploadDto) (file models.File, err error) {

	return file, nil
}
