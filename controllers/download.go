package controllers

import (
	"net/http"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Download struct {
}

func (x *Download) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

func (x *Download) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadKey [get]
func (x *Download) View(ctx context.Requester) (int, interface{}) {
	var result []byte
	//var filename string
	key, err := husk.ParseKey(ctx.FindParam("uploadKey"))

	if err != nil {

		return http.StatusInternalServerError, err
	}

	result, _, err = core.GetUploadFile(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, result
}
