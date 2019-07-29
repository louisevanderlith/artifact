package routers

import (
	"github.com/louisevanderlith/artifact/controllers"
	"github.com/louisevanderlith/droxolite"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Upload
	uplCtrl := &controllers.UploadController{}
	uplGroup := droxolite.NewRouteGroup("upload", uplCtrl)
	uplGroup.AddRoute("/", "POST", roletype.Owner, uplCtrl.Post)
	uplGroup.AddRoute("/{uploadKey:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, uplCtrl.GetByID)
	uplGroup.AddRoute("/{uploadKey:[0-9]+\x60[0-9]+}", "DELETE", roletype.Admin, uplCtrl.Delete)
	uplGroup.AddRoute("/file/{uploadKey:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, uplCtrl.GetFileBytes)
	uplGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, uplCtrl.Get)
	poxy.AddGroup(uplGroup)
	/*
		ctrlmap := EnableFilters(s, host)
		uplCtrl := controllers.NewUploadCtrl(ctrlmap)

		beego.Router("/v1/upload", uplCtrl, "post:Post")
		beego.Router("/v1/upload/:uploadKey", uplCtrl, "get:GetByID;delete:Delete")
		beego.Router("/v1/upload/all/:pagesize", uplCtrl, "get:Get")
		beego.Router("/v1/upload/file/:uploadKey", uplCtrl, "get:GetFileBytes")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["DELETE"] = roletype.Admin

	ctrlmap.Add("/v1/upload", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
