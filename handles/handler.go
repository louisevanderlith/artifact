package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong/middle"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := middle.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	view := ins.Middleware("artifact.uploads.view", scrt, ViewUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := ins.Middleware("artifact.uploads.create", scrt, CreateUpload)
	r.HandleFunc("/upload", create).Methods(http.MethodPost)

	delete := ins.Middleware("artifact.uploads.delete", scrt, DeleteUpload)
	r.HandleFunc("/upload/{key:[0-9]+\\x60[0-9]+}", delete).Methods(http.MethodDelete)

	search := ins.Middleware("artifact.uploads.search", scrt, SearchUploads)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}", search).Methods(http.MethodGet)
	r.HandleFunc("/upload/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", search).Methods(http.MethodGet)

	r.HandleFunc("/download/{key:[0-9]+`[0-9]+}", Download).Methods(http.MethodGet)

	lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "artifact.uploads.create", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodDelete,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
