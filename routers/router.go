/*
 * @Author: SingleBiu
 * @Date: 2024-10-29 15:57:44
 * @LastEditors: SingleBiu
 * @LastEditTime: 2024-10-29 16:06:41
 * @Description: file content
 */
package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mc", &controllers.McController{})
	beego.Router("/downloadfile1", &controllers.DownloadController{})
	beego.Router("/uploadfile",&controllers.UploadController{})
}
