package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/artifact/routers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()
	//r.Use(cors)
	host := os.Getenv("HOST")
	authority := "https://oauth2." + host
	provider, err := oidc.NewProvider(context.Background(), authority)
	if err != nil {
		panic(err)
	}

	r.GET("/upload/:key", upload.View)
	r.POST("/upload", upload.Create)
	r.PUT("/upload/:key", upload.Update)
	r.DELETE("/upload/:key", upload.Delete)
	r.GET("/upload/:key", upload.View)

	r.GET("/uploads", upload.Get)
	r.GET("/uploads/:pagesize/*hash", upload.Search)

	r.GET("/download/:key", controllers.Download)
	err := r.Run(":8082")

	if err != nil {
		panic(err)
	}
}
