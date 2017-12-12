package controllers

import (
	"github.com/astaxie/beego"
)

type GeneralExtendedController struct {
	beego.Controller
}

func (this *GeneralExtendedController) Prepare() {
	this.Layout = "layout-general.tpl"
}