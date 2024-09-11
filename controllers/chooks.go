package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"chookpedia_go/models"
	"log"
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



type chookResponse struct {
	Name string `json:"name"`
	Breed string `json:"breed"`
	Colour string `json:"colour"`
	PhotoURL string `json:"photo_url"`
}

type JSONChookController struct {
	beego.Controller
}

func (c *JSONChookController) Get() {
	name := c.Ctx.Input.Param(":chook")

	var chook models.Chook

	o := orm.NewOrm()
	err := o.QueryTable("chook").Filter("name", name).One(&chook)

	if err == orm.ErrMultiRows {
		// Have multiple records
		log.Fatal("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// No result
		log.Fatal("No row found")
	}

	o.Read(chook.Breed)
	log.Println(chook.Breed)

	response := chookResponse{
		Name: chook.Name,
		Colour: chook.Colour,
		Breed: chook.Breed.Name,
		PhotoURL: chook.PhotoURL,
	}

	c.Data["json"] = &response
	c.ServeJSON()
}
