package main

import (
	_ "chookpedia_go/routers"
	"chookpedia_go/models"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.InitModel()
	beego.Run()
}

