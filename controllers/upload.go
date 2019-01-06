package controllers

import (
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
		c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('failed')</script>")
	}
	defer f.Close()
	err = c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('failed')</script>")
	}
	result := res{"success", 200}
	c.Data["json"] = &result
	// c.ServeJSON()
	c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('success')</script>")
}
