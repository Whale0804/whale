package entity

type FileAddDto struct {
	Id     string `form:"id"`
	Guid   string `form:"guid"`
	Name   string `form:"name"`
	Ext    string `form:"ext"`
	Size   int    `form:"size"`
	Chunks int    `form:"chunks"`
	Chunk  int    `form:"chunk"`
	Md5    string `form:"md5"`
	Path   string `form:"path"`
}
type FinishUploadDto struct {
	Id       string `form:"id"`
	Guid     string `form:"guid"`
	Name     string `form:"name"`
	Chunks   int    `form:"chunks"`
	Filehash string `form:"filehash"`
	Size     int    `form:"size"`
	Type     string `form:"type"`
	Ext      string `form:"ext"`
	CurrPath string `form:"currPath"`
}
type CheckFile struct {
	Md5 string `form:"md5"`
}
