package controllers

import (
	"github.com/astaxie/beego"
)

type McController struct {
	beego.Controller
}

func (c *McController) Get() {
	c.TplName = "mc.html"
}
