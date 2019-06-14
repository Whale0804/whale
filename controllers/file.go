package controllers

import (
	"fmt"
	"github.com/githinkcn/whale/entity"
	"github.com/githinkcn/whale/utils"
	"io/ioutil"
	"os"
)

// Operations about Login
type FileController struct {
	BaseController
}

// @Title register
// @Description Register current logged in user session
// @Success 200 {string} logout success
// @router /upload [post]
func (this *FileController) Upload() {
	f, h, _ := this.GetFile("file")
	name := this.GetString("name")
	fileDto := entity.FileAddDto{}
	if err := this.ParseForm(&fileDto); err != nil {
		fmt.Printf("解析FormData失败：%s", err)
	}
	fmt.Println(fileDto)
	//exist, _ :=utils.PathExists("")
	fmt.Println(name)
	//得到文件的名称
	fileName := h.Filename
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer f.Close()
	if utils.FileExists("D:/whale/" + fileName) {
		fp, err := os.OpenFile("D:/whale/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		defer fp.Close()
		if err != nil {

		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fp.Write(b)
	} else {
		this.SaveToFile("file", "D:/whale/"+fileName)
	}
	this.Resp(0, "success", map[string]interface{}{})
}
