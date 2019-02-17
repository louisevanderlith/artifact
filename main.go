package main

import (
	"log"

	"github.com/louisevanderlith/artifact/routers"
	_ "github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
)

func main() {
	mode := beego.BConfig.RunMode

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
