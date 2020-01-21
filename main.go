package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/artifact/controllers"
	"github.com/louisevanderlith/artifact/controllers/upload"
	"github.com/louisevanderlith/artifact/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()
	//r.Use(cors)
	//host := os.Getenv("HOST")
	//authority := "https://oauth2." + host
	//provider, err := oidc.NewProvider(context.Background(), authority)
	//if err != nil {
	//	panic(err)
	//}

	r.GET("/upload/:key", upload.View)

	r.POST("/upload", upload.Create)
	r.PUT("/upload/:key", upload.Update)
	r.DELETE("/upload/:key", upload.Delete)
	r.GET("/upload/:key", upload.View)

	r.GET("/uploads", upload.Get)
	r.GET("/uploads/:pagesize/*hash", upload.Search)

	r.GET("/download/:key", controllers.Download)
	err := r.Run(":8082")

	/*
		r.GET("/article/:key", article.View)

			authed := r.Group("/article")
			authed.Use(droxo.Authorize())
			authed.POST("", article.Create)
			authed.PUT("/:key", article.Update)
			authed.DELETE("/:key", article.Delete)

			r.GET("/articles", article.Get)
			r.GET("/articles/:pagesize/*hash", article.Search)
	*/

	if err != nil {
		panic(err)
	}
}
