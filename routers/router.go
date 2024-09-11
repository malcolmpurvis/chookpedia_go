package routers

import (
	"chookpedia_go/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

    // @GetMapping("/")
    // fun findAll() = "Hello World"

    // @GetMapping("/blackie")
    // fun findBlackie() = "Blackie the Chook"


    // // Future endpoints:
    // // - / - index to all the chooks
    // // - /blackie - Page about blackie the chook
    // // - /chook/{name} - Page of details of a particular chook
    // // - /rest/v1/chook/{name} - JSON dump of the details of a chook
    // // - 


func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/blackie", &controllers.BlackieController{})
	beego.Router("/chook/:chook:string", &controllers.ChookController{})
	beego.Router("/rest/v1/chook/:chook:string", &controllers.JSONChookController{})
}
