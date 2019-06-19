package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/githinkcn/whale/entity"
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
	f, _, err := this.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	fileDto := entity.FileAddDto{}
	if err := this.ParseForm(&fileDto); err != nil {
		fmt.Printf("解析FormData失败：%s", err)
	}
	filePath := beego.AppConfig.String("whale_path") + fileDto.Id + "/"
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
	dto := &entity.FinishUploadDto{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &dto)
	fmt.Println(dto)
	ab, _ := os.Create(beego.AppConfig.String("whale_path") + dto.Name)
	for i := 0; i < dto.Chunks; i++ {
		f, _ := os.OpenFile(beego.AppConfig.String("whale_path")+dto.Id+"/"+strconv.Itoa(i), os.O_RDONLY, os.ModePerm)
		b, _ := ioutil.ReadAll(f)
		ab.Write(b)
		f.Close()
		os.Remove(beego.AppConfig.String("whale_path") + dto.Id + "/" + strconv.Itoa(i))
	}
	os.RemoveAll(beego.AppConfig.String("whale_path") + dto.Id + "/")
	defer ab.Close()
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
