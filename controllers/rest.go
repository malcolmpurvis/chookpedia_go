package controllers

import (
	"chookpedia_go/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type chookResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Colour   string `json:"colour"`
	Breed    int    `json:"breed"`
	PhotoURL string `json:"photo_url"`
}

type JSONChookController struct {
	beego.Controller
}

func (c *JSONChookController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	chook, err := models.FindChook(id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	response := chookResponse{
		Id:       chook.Id,
		Name:     chook.Name,
		Colour:   chook.Colour,
		Breed:    chook.Breed.Id,
		PhotoURL: chook.PhotoURL,
	}

	c.Data["json"] = &response
	c.ServeJSON()
}

type chookCollectionResponse []chookResponse

type JSONChookCollectionController struct {
	beego.Controller
}

func (c *JSONChookCollectionController) Get() {
	chooks, err := models.AllChooks()

	if err != nil {
		c.CustomAbort(404, err.Error())
	}

	response := chookCollectionResponse{}

	for _, chook := range chooks {
		response = append(response,
			chookResponse{
				Id:       chook.Id,
				Name:     chook.Name,
				Colour:   chook.Colour,
				Breed:    chook.Breed.Id,
				PhotoURL: chook.PhotoURL,
			})
	}

	c.Data["json"] = &response
	c.ServeJSON()
}

type breedResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

type JSONBreedController struct {
	beego.Controller
}

func (c *JSONBreedController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	breed, err := models.FindBreed(id)

	if err == orm.ErrNoRows {
		c.CustomAbort(404, "No primary key found.")
	} else if err == orm.ErrMissPK {
		c.CustomAbort(500, "No primary key found.")
	}

	response := breedResponse{
		Id:     breed.Id,
		Name:   breed.Name,
		Origin: breed.Origin,
	}

	c.Data["json"] = &response
	c.ServeJSON()
}

type breedCollectionResponse []breedResponse

type JSONBreedCollectionController struct {
	beego.Controller
}

func (c *JSONBreedCollectionController) Get() {
	var breeds []*models.Breed

	o := orm.NewOrm()
	_, err := o.QueryTable("breed").All(&breeds)

	if err != nil {
		c.CustomAbort(404, err.Error())
	}

	response := breedCollectionResponse{}

	for _, breed := range breeds {
		response = append(response,
			breedResponse{
				Id:     breed.Id,
				Name:   breed.Name,
				Origin: breed.Origin,
			})
	}

	c.Data["json"] = &response
	c.ServeJSON()
}
