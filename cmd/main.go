package main

import (
	"flag"
	"github.com/louisevanderlith/artifact/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/artifact/core"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "artifact", "Token target 'aud'")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8082",
		Handler:      handles.SetupRoutes(*issuer, *audience),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
