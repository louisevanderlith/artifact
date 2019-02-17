// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/artifact/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango/enums"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)
	uplCtrl := controllers.NewUploadCtrl(ctrlmap)

	beego.Router("/v1/upload", uplCtrl, "post:Post")
	beego.Router("/v1/upload/:uploadKey", uplCtrl, "get:GetByID")
	beego.Router("/v1/upload/all/:pagesize", uplCtrl, "get:Get")
	beego.Router("/v1/upload/file/:uploadKey", uplCtrl, "get:GetFileBytes")
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Owner

	ctrlmap.Add("/upload", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
