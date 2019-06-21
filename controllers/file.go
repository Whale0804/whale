package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/githinkcn/whale/common"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/service"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Operations about Login
type FileController struct {
	BaseController
}

// @Title 文件上传
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /upload [post]
func (this *FileController) Upload() {
	user, _ := this.GetCurrentUser()
	f, _, err := this.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	fileDto := entity.FileAddDto{}
	if err := this.ParseForm(&fileDto); err != nil {
		fmt.Printf("解析FormData失败：%s", err)
	}
	fileService := service.FileService{}
	folder, _ := fileService.FindTopFolderByUserId(user.Id)
	filePath := beego.AppConfig.String("whale_path") + folder.FolderName + "/" + fileDto.Id + "/"
	fileDto.Path = filePath
	err1 := os.MkdirAll(filePath, os.ModePerm)
	if err1 != nil {
		fmt.Println(err1)
	}

	this.SaveToFile("file", filePath+strconv.Itoa(fileDto.Chunk))
	this.Resp(0, "success", map[string]interface{}{
		"data": fileDto,
	})
}

// @Title 通知合并
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /finish [post]
func (this *FileController) UploadFinish() {
	user, _ := this.GetCurrentUser()
	fileService := service.FileService{}
	folder, _ := fileService.FindTopFolderByUserId(user.Id)
	dto := &entity.FinishUploadDto{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dto)
	ab, _ := os.Create(beego.AppConfig.String("whale_path") + folder.FolderName + "/" + dto.Name)
	for i := 0; i < dto.Chunks; i++ {
		f, _ := os.OpenFile(beego.AppConfig.String("whale_path")+folder.FolderName+"/"+dto.Id+"/"+strconv.Itoa(i), os.O_RDONLY, os.ModePerm)
		b, _ := ioutil.ReadAll(f)
		ab.Write(b)
		f.Close()
		os.Remove(beego.AppConfig.String("whale_path") + folder.FolderName + "/" + dto.Id + "/" + strconv.Itoa(i))
	}
	os.RemoveAll(beego.AppConfig.String("whale_path") + folder.FolderName + "/" + dto.Id + "/")
	defer ab.Close()
	//存储文件信息
	fileService.SaveFile(dto)
	this.Resp(0, "success", map[string]interface{}{})
}

// @Title 秒传
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /check [post]
func (this *FileController) CheckFile() {
	dto := &entity.CheckFile{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dto)
	if strings.EqualFold(dto.Md5, "46dc080ca248abcad701351964fe8b45") {
		this.Resp(0, "success", map[string]interface{}{
			"isExist": true,
		})
	} else {
		this.Resp(0, "success", map[string]interface{}{
			"isExist": false,
		})
	}
}

// @Title 创建文件夹
// @Description 创建文件夹
// @Success 200 {string} logout success
// @router /folder [post]
func (this *FileController) CreateFolder() {
	this.GetString("folder_name")
	fmt.Println(1)
	user, _ := this.GetCurrentUser()
	//查询用户关联的
	fileService := service.FileService{}
	folder, err := fileService.FindTopFolderByUserId(user.Id)
	if err != nil {
		this.Fail(common.ErrCreateFolder)
		return
	}
	FolderPath := beego.AppConfig.String("whale_path") + folder.FolderName + "/"
	fmt.Println(FolderPath)
}
