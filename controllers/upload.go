package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

type fileObj struct {
	Name string
	Data []byte
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
	// GetFiles return multi-upload files
	files, err := c.GetFiles("saveFileByForm")
	if err != nil {
		c.Ctx.WriteString("GetFiles failed")
	}
	for i, v := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			c.Ctx.WriteString(v.Filename + " upload failed\r\n")
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create("./files/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			c.Ctx.WriteString(v.Filename + " upload failed\r\n")
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			c.Ctx.WriteString(v.Filename + " upload failed\r\n")
		}
		c.Ctx.WriteString(v.Filename + " upload success\r\n")
	}
}

// SaveFileByFormNoFresh SaveFileByFormNoFresh
func (c *UploadController) SaveFileByFormNoFresh() {
	var scriptsStr string
	files, err := c.GetFiles("saveFileByForm")
	if err != nil {
		c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('GetFiles failed', 'error')</script>")
	}
	for i, v := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			scriptsStr += v.Filename + " upload failed "
			continue
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create("./files/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			scriptsStr += v.Filename + " upload failed "
			continue
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			scriptsStr += v.Filename + " upload failed "
			continue
		}
		scriptsStr += v.Filename + " upload success "
	}
	c.Ctx.WriteString("<script type='text/javascript'>parent.showRes('" + scriptsStr + "', 'success')</script>")
}

// SaveFileByAjaxForm SaveFileByAjaxForm
func (c *UploadController) SaveFileByAjaxForm() {
	var resStr string
	files, err := c.GetFiles("saveFileByForm")
	if err != nil {
		str, _ := json.Marshal(res{"GetFile falied", 420})
		c.Ctx.WriteString(string(str[:]))
	}
	for i, v := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			resStr += v.Filename + " upload failed "
			continue
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create("./files/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			resStr += v.Filename + " upload failed "
			continue
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			resStr += v.Filename + " upload failed "
			continue
		}
		resStr += v.Filename + " upload success "
	}
	str, _ := json.Marshal(res{resStr, 200})
	c.Ctx.WriteString(string(str[:]))
}

// SaveFileByAxios SaveFileByAxios
func (c *UploadController) SaveFileByAxios() {
	var result res
	files, err := c.GetFiles("saveFileByForm")
	if err != nil {
		result = res{"GetFile falied", 420}
		c.Data["json"] = &result
		c.ServeJSON()
	}
	for i, v := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			result.Msg += v.Filename + " upload failed "
			continue
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create("./files/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			result.Msg += v.Filename + " upload failed "
			continue
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			result.Msg += v.Filename + " upload failed "
			continue
		}
		result.Msg += v.Filename + " upload success "
	}
	result.Code = 200
	c.Data["json"] = &result
	c.ServeJSON()
}

// SaveFileByFileReader SaveFileByFileReader
func (c *UploadController) SaveFileByFileReader() {
	var result res
	var filesObj fileObj
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &filesObj); err != nil {
		log.Print(err)
		result = res{"json Unmarshal falied", 420}
		c.Data["json"] = &result
		c.ServeJSON()
	}
	log.Print(filesObj)
	fileName := filesObj.Name
	dst, err := os.Create("./files/" + fileName)
	defer dst.Close()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		result.Msg += fileName + " upload failed "
	}
	//copy the uploaded file to the destination file
	if _, err := io.WriteString(dst, string(filesObj.Data)); err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		result.Msg += fileName + " upload failed "
	}
	result.Code = 200
	c.Data["json"] = &result
	c.ServeJSON()
}
