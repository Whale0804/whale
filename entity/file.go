package entity

type FileAddDto struct {
	Id     string `form:"id"`
	Name   string `form:"name"`
	Ext    string `form:"ext"`
	Size   int    `form:"size"`
	Chunks int    `form:"chunks"`
	Chunk  int    `form:"chunk"`
	Md5    string `form:"md5"`
}
type FinishUploadDto struct {
	Name   string `form:"name"`
	Chunks int    `form:"chunks"`
	Path   string `form:"path"`
}
