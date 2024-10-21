/*
 * @Author: SingleBiu
 * @Date: 2024-10-21 20:39:56
 * @LastEditors: SingleBiu
 * @LastEditTime: 2024-10-21 21:35:59
 * @Description: file content
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {
	c.Ctx.Output.Download("static/file/Picacg.apk", "Picacg.apk")
}
