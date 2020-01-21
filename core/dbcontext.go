package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Uploads husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Uploads: husk.NewTable(Upload{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Uploads.Save()
}
