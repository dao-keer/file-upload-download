package controllers

import (
	"io/ioutil"
	"log"

	"github.com/astaxie/beego"
)

// GetFilesController GetFilesController
type GetFilesController struct {
	beego.Controller
}

// GetFilesList GetFilesList
func (c *GetFilesController) GetFilesList() {
	dirList, e := ioutil.ReadDir("./static/files")
	var files []string
	if e != nil {
		log.Print(e)
		result := fileRes{"ReadDir error", 200, fileList{files}}
		c.Data["json"] = &result
		c.ServeJSON()
	}
	for _, v := range dirList {
		files = append(files, v.Name())
	}
	log.Print(files)
	result := fileRes{"ReadDir success", 200, fileList{files}}
	c.Data["json"] = &result
	c.ServeJSON()
}
