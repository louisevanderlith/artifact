package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/artifact/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilters(s, host)
	uplCtrl := controllers.NewUploadCtrl(ctrlmap)

	beego.Router("/v1/upload", uplCtrl, "post:Post")
	beego.Router("/v1/upload/:uploadKey", uplCtrl, "get:GetByID")
	beego.Router("/v1/upload/all/:pagesize", uplCtrl, "get:Get")
	beego.Router("/v1/upload/file/:uploadKey", uplCtrl, "get:GetFileBytes")
}

func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner

	ctrlmap.Add("/v1/upload", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
