package main

import (
	"flag"
	"github.com/louisevanderlith/artifact/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/artifact/core"
)

func main() {
	host := flag.String("host", "localhost", "Domain Host")
	authrty := flag.String("authority", "http://localhost:8086", "Authority Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8082",
		Handler:      handles.SetupRoutes(*host, *srcSecrt, *authrty),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
