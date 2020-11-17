package routers

import (
	"kfktools/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})

	beego.Router("/topic/list", &controllers.TopicController{})
	beego.Router("/topic/create", &controllers.TopicController{}, "*:Create")
}
