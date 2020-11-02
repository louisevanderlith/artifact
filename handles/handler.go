package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)


func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	//view := ins.Middleware("artifact.uploads.view", scrt, ViewUpload)
	r.Handle("/upload/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewUpload))).Methods(http.MethodGet)

	//create := ins.Middleware("artifact.uploads.create", scrt, CreateUpload)
	r.Handle("/upload", mw.Handler(http.HandlerFunc(CreateUpload))).Methods(http.MethodPost)

	//delete := ins.Middleware("artifact.uploads.delete", scrt, DeleteUpload)
	r.Handle("/upload/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(DeleteUpload))).Methods(http.MethodDelete)

	//search := ins.Middleware("artifact.uploads.search", scrt, SearchUploads)
	r.Handle("/upload/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchUploads))).Methods(http.MethodGet)
	r.Handle("/upload/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchUploads))).Methods(http.MethodGet)

	r.HandleFunc("/download/{key:[0-9]+`[0-9]+}", Download).Methods(http.MethodGet)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "artifact.uploads.create", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
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
