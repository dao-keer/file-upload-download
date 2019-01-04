package controllers

import (
	"github.com/astaxie/beego"
)

// MainController MainController
type MainController struct {
	beego.Controller
}

// Get Get
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
