package controllers

import (
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/artifact/logic"
	"github.com/louisevanderlith/mango/control"
)

type UploadController struct {
	control.APIController
}

func NewUploadCtrl(ctrlMap *control.ControllerMap) *UploadController {
	result := &UploadController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]core.Upload} []core.Upload
// @router /all/:pagesize [get]
func (req *UploadController) Get() {
	page, size := req.GetPageData()

	results := core.GetUploads(page, size)

	req.Serve(results, nil)
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadKey			path	husk.Key 	true		"Key of the file you require"
// @Success 200 {core.Upload} core.Upload
// @router /:uploadKey [get]
func (req *UploadController) GetByID() {
	key, err := husk.ParseKey(req.Ctx.Input.Param(":uploadKey"))

	if err != nil {
		req.Serve(nil, err)
		return
	}

	req.Serve(core.GetUpload(key))
}

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadKey [get]
func (req *UploadController) GetFileBytes() {
	var result []byte
	var filename string
	key, err := husk.ParseKey(req.Ctx.Input.Param(":uploadKey"))

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.ServeBinary([]byte(err.Error()), "")
		return
	}

	result, filename, err = core.GetUploadFile(key)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		result = []byte(err.Error())
	}

	req.ServeBinary(result, filename)
}

// @Title UploadFile
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Param	body		body 	core.Upload	true		"body for upload content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *UploadController) Post() {
	key := husk.CrazyKey()

	info := req.GetString("info")
	infoHead, err := logic.GetInfoHead(info)

	if err != nil {
		req.Serve(key, err)
		return
	}

	file, header, err := req.GetFile("file")

	if err != nil {
		req.Serve(key, err)
		return
	}

	defer file.Close()

	key, err = logic.SaveFile(file, header, infoHead)

	req.Serve(key, err)
}
