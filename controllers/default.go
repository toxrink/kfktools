package controllers

import (
	"github.com/astaxie/beego"
)

//MainController 页面跳转
type MainController struct {
	beego.Controller
}

//Get 首页
func (c *MainController) Get() {
	c.TplName = "index.html"
}
