package routers

import (
	"Movie-Picker/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/searcha", &controllers.SearchAgainController{})
	
}
