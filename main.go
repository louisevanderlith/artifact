package main

import (
	"flag"
	"github.com/louisevanderlith/artifact/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/artifact/core"
)

func main() {
	securty := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8082",
		Handler:      handles.SetupRoutes(*securty, *srcSecrt),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
