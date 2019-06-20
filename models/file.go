package models

import (
	"time"
)

//文件实体类
type File struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	Ext        string    `json:"ext" orm:"column(ext)" description:"后缀名"`
	FileName   string    `json:"filename" orm:"column(file_name)" description:"文件名"`
	FileSize   string    `json:"filesize" orm:"column(file_size)" description:"文件大小"`
	FilePath   string    `json:"filepath" orm:"column(file_path)" description:"完整文件路径"`
	FileHash   string    `json:"filehash" orm:"column(file_hash)" description:"文件hash"`
	FileType   string    `json:"filetype" orm:"column(file_type)" description:"文件类型"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

//用户文件关联实体类
type UserFile struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	UserId     int       `json:"userId" orm:"column(user_id)" description:"用户Id"`
	FileId     int       `json:"fileId" orm:"column(file_id)" description:"文件Id"`
	FileHash   string    `json:"filehash" orm:"column(file_hash)" description:"文件hash"`
	FileName   string    `json:"filename" orm:"column(file_name)" description:"文件名"`
	FileSize   string    `json:"filesize" orm:"column(file_size)" description:"文件大小"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

//文件夹实体类
type Folder struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	FolderName string    `json:"foldername" orm:"column(folder_name)" description:"文件夹名称"`
	ParentId   int       `json:"parentId" orm:"column(parent_id)" description:"文件夹父级"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

//用户文件夹实体类
type UserFolder struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	UserId     int       `json:"userId" orm:"column(user_id)" description:"用户主键"`
	FolderId   int       `json:"folderId" orm:"column(folder_id)" description:"文件夹主键"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

//用户文件夹文件实体类
type UserFolderFile struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	UserId     int       `json:"userId" orm:"column(user_id)" description:"用户主键"`
	FolderId   int       `json:"folderId" orm:"column(folder_id)" description:"文件夹主键"`
	FileId     int       `json:"fileId" orm:"column(file_id)" description:"文件Id"`
	Type       int       `json:"type" orm:"column(type);size(1)" description:"数据类型：0-文件夹，1-文件"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}
