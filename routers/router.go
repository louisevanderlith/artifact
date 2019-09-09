package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/artifact/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Owner, mix.JSON, &controllers.Upload{})
	e.JoinPath(e.Router().(*mux.Router), "/download/{key:[0-9]+\x60[0-9]+}", "Download File", http.MethodGet, roletype.Unknown, mix.Octet, controllers.Download)
}
