package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/githinkcn/whale/entity"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	filePath := "./whale/" + fileDto.Id + "/"
	err1 := os.MkdirAll(filePath, os.ModePerm)
	if err1 != nil {
		fmt.Println(err1)
	}
	this.SaveToFile("file", filePath+strconv.Itoa(fileDto.Chunk))
	this.Resp(0, "success", map[string]interface{}{
		"data":     fileDto,
		"filePath": filePath,
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
	ab, err := os.Create("./whale/" + dto.Name)
	fmt.Println(err)
	for i := 0; i < dto.Chunks; i++ {
		f, _ := os.OpenFile(dto.Path+strconv.Itoa(i), os.O_RDONLY, os.ModePerm)
		b, _ := ioutil.ReadAll(f)
		ab.Write(b)
		f.Close()
		os.Remove(dto.Path + strconv.Itoa(i))
	}
	os.RemoveAll(dto.Path)
	defer ab.Close()
	this.Resp(0, "success", map[string]interface{}{})
}

// @Title 秒传
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /check [post]
func (this *FileController) CheckFile() {
	fmt.Println(this.GetString("md5"))
	this.Resp(0, "success", map[string]interface{}{})
}
