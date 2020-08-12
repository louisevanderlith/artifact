package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware(http.DefaultClient, "artifact.uploads.view", scrt, securityUrl, managerUrl, ViewUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := kong.ResourceMiddleware(http.DefaultClient, "artifact.uploads.create", scrt, securityUrl, managerUrl, CreateUpload)
	r.HandleFunc("/upload", create).Methods(http.MethodPost)

	delete := kong.ResourceMiddleware(http.DefaultClient, "artifact.uploads.delete", scrt, securityUrl, managerUrl, DeleteUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", delete).Methods(http.MethodDelete)

	search := kong.ResourceMiddleware(http.DefaultClient, "artifact.uploads.search", scrt, securityUrl, managerUrl, SearchUploads)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}", search).Methods(http.MethodGet)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", search).Methods(http.MethodGet)

	dwnld := kong.ResourceMiddleware(http.DefaultClient, "artifact.download", scrt, securityUrl, managerUrl, Download)
	r.HandleFunc("/download/{key:[0-9]+`[0-9]+}", dwnld).Methods(http.MethodGet)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "artifact.download", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
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
