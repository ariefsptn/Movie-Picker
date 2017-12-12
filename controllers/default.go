package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "Movies Picker!"
	
	this.TplName = "index.tpl"
	
}


