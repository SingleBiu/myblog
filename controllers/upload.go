/*
 * @Author: SingleBiu
 * @Date: 2024-10-29 15:55:45
 * @LastEditors: SingleBiu
 * @LastEditTime: 2024-10-29 16:48:27
 * @Description: file content
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
 
	c.TplName="index.html"
}

func (c *UploadController) Post() {
	file, information, err := c.GetFile("file")  //返回文件，文件信息头，错误信息
    if err != nil {
        c.Ctx.WriteString("File retrieval failure")
        return
    }
    defer file.Close()    //关闭上传的文件，否则出现临时文件不清除的情况  mmp错了好多次啊

    filename := information.Filename           //将文件信息头的信息赋值给filename变量
    err = c.SaveToFile("file", "static/upload"+filename) //保存文件的路径。保存在static/upload中   （文件名）
    if err != nil {
        c.Ctx.WriteString("File upload failed!")
    } else {
        c.Ctx.WriteString("File upload succeed!")  //上传成功后显示信息
    }
    c.TplName = "index.html"             //停留在当前界面
}