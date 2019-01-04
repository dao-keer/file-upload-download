package routers

import (
	"file-upload-download/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/saveFileByForm", &controllers.UploadController{}, "post:SaveFileByForm")
}
