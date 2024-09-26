package controllers

import (
	"chookpedia_go/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type BlackieController struct {
	beego.Controller
}

func (c *BlackieController) Get() {
	id := 1

	chook, err := models.FindChook(id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	breed, err := models.FindBreed(chook.Breed.Id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	c.Data["Name"] = chook.Name
	c.Data["Colour"] = chook.Colour
	c.Data["Breed"] = breed.Name
	c.TplName = "chook.tpl"
}

type ChookController struct {
	beego.Controller
}

func (c *ChookController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	chook, err := models.FindChook(id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	breed, err := models.FindBreed(chook.Breed.Id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	c.Data["Name"] = chook.Name
	c.Data["Colour"] = chook.Colour
	c.Data["Breed"] = breed.Name
	c.TplName = "chook.tpl"
}
