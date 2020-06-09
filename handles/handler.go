package handles

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

func SetupRoutes(host, authUrl, scrt string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware("artifact.uploads.view", scrt, authUrl, ViewUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("artifact.uploads.create", scrt, authUrl, CreateUpload)
	r.HandleFunc("/upload", create).Methods(http.MethodPost)

	delete := kong.ResourceMiddleware("artifact.uploads.delete", scrt, authUrl, DeleteUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", delete).Methods(http.MethodDelete)

	search := kong.ResourceMiddleware("artifact.uploads.search", scrt, authUrl, SearchUploads)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}", search).Methods(http.MethodGet)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", search).Methods(http.MethodGet)

	dwnld := kong.ResourceMiddleware("artifact.download", scrt, authUrl, Download)
	r.HandleFunc("/download/{key:[0-9]+`[0-9]+}", dwnld).Methods(http.MethodGet)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
