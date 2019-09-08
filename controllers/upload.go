package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/artifact/logic"
)

type Upload struct {
}

func (req *Upload) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetUploads(1, 10)

	return http.StatusOK, results
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]core.Upload} []core.Upload
// @router /all/:pagesize [get]
func (req *Upload) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()

	results := core.GetUploads(page, size)

	return http.StatusOK, results
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadKey			path	husk.Key 	true		"Key of the file you require"
// @Success 200 {core.Upload} core.Upload
// @router /:uploadKey [get]
func (req *Upload) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("uploadKey"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	record, err := core.GetUpload(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, record
}

// @Title UploadFile
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Param	body		body 	core.Upload	true		"body for upload content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Upload) Create(ctx context.Requester) (int, interface{}) {
	file, header, err := ctx.File("file")

	if err != nil {
		return http.StatusBadRequest, err
	}

	defer file.Close()

	info := ctx.FindFormValue("info")
	infoHead, err := logic.GetInfoHead(info)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	log.Printf("Size: %v\tKey: %s\r", header.Size, infoHead.ItemKey.String())

	var b bytes.Buffer
	copied, err := io.Copy(&b, file)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	key, err := logic.SaveFile(b.Bytes(), copied, header, infoHead)

	if err != nil {
		panic(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, key
}

func (req *Upload) Update(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// @router /:uploadKey [delte]
func (req *Upload) Delete(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("uploadKey"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = core.RemoveUpload(key)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Completed"
}
