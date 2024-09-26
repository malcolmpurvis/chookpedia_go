package controllers

import (
	"chookpedia_go/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()

	chooks, err := models.AllChooks()

	if err != nil {
		c.CustomAbort(404, err.Error())
	}

	for _, chook := range chooks {
		o.Read(chook.Breed)
	}

	c.Data["chooks"] = chooks
	c.TplName = "all_chooks.tpl"
}
