package routers

import (
	"github.com/louisevanderlith/artifact/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Owner, mix.JSON, &controllers.Upload{})
	e.JoinBundle("/", roletype.Unknown, mix.Octet, &controllers.Download{})
}
