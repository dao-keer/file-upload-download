package controllers

import (
	"github.com/astaxie/beego"
)

// MainController MainController
type MainController struct {
	beego.Controller
}

// Res Res
type res struct {
	Msg  string
	Code int
}

type fileObj struct {
	Name string
	Data []byte
}

type fileRes struct {
	Msg  string
	Code int
	Data fileList
}

type fileList struct {
	FilesList []string
}

// Get Get
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
