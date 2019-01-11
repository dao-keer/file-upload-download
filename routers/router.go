package routers

import (
	"file-upload-download/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/getFilesList", &controllers.GetFilesController{}, "get:GetFilesList")
	beego.Router("/api/saveFileByForm", &controllers.UploadController{}, "post:SaveFileByForm")
	beego.Router("/api/saveFileByFormNoFresh", &controllers.UploadController{}, "post:SaveFileByFormNoFresh")
	beego.Router("/api/saveFileByAxios", &controllers.UploadController{}, "post:SaveFileByAxios")
	beego.Router("/api/saveFileByAjaxForm", &controllers.UploadController{}, "post:SaveFileByAjaxForm")
	beego.Router("/api/saveFileByFileReader", &controllers.UploadController{}, "post:SaveFileByFileReader")
	beego.Router("/api/getFile", &controllers.DownloadController{}, "get:GetFileByGet")
}
