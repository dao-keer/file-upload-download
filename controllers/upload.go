package controllers

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

// UploadController UploadController
type UploadController struct {
	beego.Controller
}

// Res Res
type res struct {
	Msg  string
	Code int
}

// IEVersion IEVersion
func IEVersion(userAgent string) string {
	isIE := false
	isEdge := false
	isIE11 := false
	hasCompatible := strings.Index(userAgent, "compatible")
	hasMSIE := strings.Index(userAgent, "MSIE")
	hasEdge := strings.Index(userAgent, "Edge")
	hasTrident := strings.Index(userAgent, "Trident")
	hasRV := strings.Index(userAgent, "rv:11.0")
	if hasCompatible > -1 && hasMSIE > -1 {
		isIE = true
	}
	if hasEdge > -1 && !isIE {
		isEdge = true
	}
	if hasTrident > -1 && hasRV > -1 {
		isIE11 = true
	}
	if isIE {
		re := regexp.MustCompile("MSIE (\\d+\\.\\d+);")
		log.Print(re.FindAllString("paranormal", -1))
		fIEVersion := 7
		if fIEVersion == 7 {
			return "7"
		} else if fIEVersion == 8 {
			return "8"
		} else if fIEVersion == 9 {
			return "9"
		} else if fIEVersion == 10 {
			return "10"
		} else {
			return "6" //IE版本<=7
		}
	} else if isEdge {
		return "edge"
	} else if isIE11 {
		return "11" //IE11
	} else {
		return "0" //不是ie浏览器
	}
}

// SaveFileByForm SaveFileByForm
func (c *UploadController) SaveFileByForm() {
	f, h, err := c.GetFile("saveFileByForm")
	if err != nil {
		c.Ctx.WriteString("GetFile failed")
	}
	defer f.Close()
	err = c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		c.Ctx.WriteString("SaveToFile failed")
	}
	c.Ctx.WriteString("上传成功")
}

// SaveFileByFormNoFresh SaveFileByFormNoFresh
func (c *UploadController) SaveFileByFormNoFresh() {
	f, h, err := c.GetFile("saveFileByForm")
	if err != nil {
		c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('获取上传文件失败', 'error')</script>")
	}
	defer f.Close()
	err = c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('保存文件失败')</script>")
	}
	c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('上传文件成功', 'success')</script>")
}

// SaveFileByAjaxForm SaveFileByAjaxForm
func (c *UploadController) SaveFileByAjaxForm() {
	log.Print(c.Ctx.Request.Header.Get("user-agent"))
	f, h, err := c.GetFile("saveFileByForm")
	if err != nil {
		str, _ := json.Marshal(res{"GetFile falied", 420})
		c.Ctx.WriteString(string(str[:]))
	}
	defer f.Close()
	err = c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		str, _ := json.Marshal(res{"SaveToFile falied", 430})
		c.Ctx.WriteString(string(str[:]))
	}
	str, _ := json.Marshal(res{"success", 200})
	c.Ctx.WriteString(string(str[:]))
}

// SaveFileByAxios SaveFileByAxios
func (c *UploadController) SaveFileByAxios() {
	f, h, err := c.GetFile("saveFileByForm")
	if err != nil {
		result := res{"GetFile falied", 420}
		c.Data["json"] = &result
		c.ServeJSON()
	}
	defer f.Close()
	err = c.SaveToFile("saveFileByForm", "./files/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		result := res{"SaveToFile falied", 430}
		c.Data["json"] = &result
		c.ServeJSON()
	}
	result := res{"success", 200}
	c.Data["json"] = &result
	c.ServeJSON()
}
