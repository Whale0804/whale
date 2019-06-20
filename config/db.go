package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/githinkcn/whale/models"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER = "mysql"
)

func InitDB() {
	username := beego.AppConfig.String("mysql_user")
	password := beego.AppConfig.String("mysql_password")
	url := beego.AppConfig.String("mysql_url")
	db := beego.AppConfig.String("mysql_db")
	orm.RegisterDriver(DB_DRIVER, orm.DRMySQL)
	orm.RegisterModelWithPrefix("whale_",
		new(models.User),
		new(models.File),
		new(models.UserFile),
		new(models.Folder),
		new(models.UserFolder),
		new(models.UserFolderFile))
	orm.RegisterDataBase("default", DB_DRIVER,
		username+":"+password+"@tcp("+url+")/"+db+"?charset=utf8", 30)
}

const (
	USER_TABLE             = "whale_user"
	FILE_TABLE             = "whale_file"
	USER_FILE_TABLE        = "whale_user_file"
	FOLDER_TABLE           = "whale_folder"
	USER_FOLDER_TABLE      = "whale_user_folder"
	USER_FOLDER_FILE_TABLE = "whale_user_folder_file"
)
