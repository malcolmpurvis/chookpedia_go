package main

import (
	"chookpedia_go/models"
	_ "chookpedia_go/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.InitModel()
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
