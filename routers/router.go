package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mc", &controllers.McController{})
	beego.Router("/downloadfile1", &controllers.DownloadController{})
}
