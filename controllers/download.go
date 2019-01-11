package controllers

import (
	"github.com/astaxie/beego"
)

// DownloadController DownloadController
type DownloadController struct {
	beego.Controller
}

// GetFileByGet GetFileByGet
func (c *DownloadController) GetFileByGet() {
	fileName := c.GetString("FileName")
	c.Ctx.Output.Download("static/files/"+fileName, fileName)
}
