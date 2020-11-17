package controllers

import (
	"kfktools/utils"

	"github.com/astaxie/beego"
)

//TopicController 页面跳转
type TopicController struct {
	beego.Controller
}

//Get 获取主题列表
func (c *TopicController) Get() {
	json := utils.GetKafkaTopic()
	c.Data["json"] = &json
	c.ServeJSON()
}

//Create 创建主题
func (c *TopicController) Create() {
	json := utils.GetKafkaTopic()
	c.Data["json"] = &json
	c.ServeJSON()
}
