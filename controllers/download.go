package controllers

import (
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
func Download(ctx context.Requester) (int, interface{}) {
	var result []byte
	//var filename string
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, nil
	}

	result, _, err = core.GetUploadFile(key)

	if err != nil {
		log.Println(err)
		return http.StatusNotFound, nil
	}

	return http.StatusOK, result
}
