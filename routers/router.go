package routers

import (
	"chookpedia_go/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

// // Future endpoints:
// // - / - index to all the chooks
// // - /blackie - Page about blackie the chook
// // - /chook/{name} - Page of details of a particular chook
// // - /rest/v1/chook/{name} - JSON dump of the details of a chook
// // -

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/blackie", &controllers.BlackieController{})
	beego.Router("/chook/:id:int", &controllers.ChookController{})

	// REST endpoints

	beego.Router("/rest/v1/chook/", &controllers.JSONChookCollectionController{})
	beego.Router("/rest/v1/chook/:id:int", &controllers.JSONChookController{})
	beego.Router("/rest/v1/breed/", &controllers.JSONBreedCollectionController{})
	beego.Router("/rest/v1/breed/:id:int", &controllers.JSONBreedController{})
}
