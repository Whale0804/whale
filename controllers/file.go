package controllers

import (
	"fmt"
	"strings"
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
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	fmt.Println("文件名称:")
	fmt.Println(fileName)
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
}
