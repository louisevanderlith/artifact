package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Uploads husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Uploads: husk.NewTable(Upload{}),
	}
}

func Shutdown() {
	ctx.Uploads.Save()
}
