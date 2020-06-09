package handles

import (
	"bytes"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadKey [get]
func Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	result, file, err := core.GetUploadFile(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.Octet(file, bytes.NewReader(result)))

	if err != nil {
		log.Println(err)
		return
	}
}
