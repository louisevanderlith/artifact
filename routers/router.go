package routers

import (
	"github.com/louisevanderlith/artifact/controllers"
	"github.com/louisevanderlith/droxolite/mix"

	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Upload
	uplCtrl := &controllers.UploadController{}
	uplGroup := routing.NewRouteGroup("upload", mix.JSON)
	uplGroup.AddRoute("Create Upload", "", "POST", roletype.Owner, uplCtrl.Post)
	uplGroup.AddRoute("Get Upload By ID", "/{uploadKey:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, uplCtrl.GetByID)
	uplGroup.AddRoute("Delete Upload", "/{uploadKey:[0-9]+\x60[0-9]+}", "DELETE", roletype.Admin, uplCtrl.Delete)
	uplGroup.AddRoute("Get All Uploads", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, uplCtrl.Get)
	poxy.AddGroup(uplGroup)

	downGroup := routing.NewRouteGroup("download", mix.Octet)
	downGroup.AddRoute("Get Upload Raw", "/{uploadKey:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, controllers.GetFileBytes)
}
