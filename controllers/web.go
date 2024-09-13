package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BlackieController struct {
	beego.Controller
}

func (c *BlackieController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type ChookController struct {
	beego.Controller
}

func (c *ChookController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
