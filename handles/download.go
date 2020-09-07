package handles

import (
	"bytes"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"

	"github.com/louisevanderlith/artifact/core"
)

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadKey [get]
func Download(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

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

	err = mix.Write(w, mix.Octet(file, bytes.NewReader(result)))

	if err != nil {
		log.Println(err)
		return
	}
}
