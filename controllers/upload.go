package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

// UploadController UploadController
type UploadController struct {
	beego.Controller
}

// baseRouter implemented global settings for all other routers.
type res struct {
	msg  string
	code int
}

// SaveFileByForm SaveFileByForm
func (c *UploadController) SaveFileByForm() {
	f, h, err := c.GetFile("saveFileByForm")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	result := res{"success", 200}
	c.Data["json"] = &result
	c.ServeJSON()
}
