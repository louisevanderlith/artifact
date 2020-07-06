package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
)

func GetUploads(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetUploads(1, 10)

	if err != nil {
		log.Println("GetUploads Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]core.Upload} []core.Upload
// @router /all/:pagesize [get]
func SearchUploads(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()

	results, err := core.GetUploads(page, size)

	if err != nil {
		log.Println("GetUploads Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// ViewUpload
func ViewUpload(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	record, err := core.GetUpload(key)

	if err != nil {
		log.Println("GetUpload", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(record))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// @Title UploadFile
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Param	body		body 	core.Upload	true		"body for upload content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateUpload(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	file, header, err := ctx.File("file")

	if err != nil {
		log.Println("File Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer file.Close()

	info := ctx.FindFormValue("info")
	infoHead, err := core.GetInfoHead(info)

	if err != nil {
		log.Println("GetInfoHead Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	key, err := core.SaveFile(file, header, infoHead)

	if err != nil {
		log.Println("SaveFile Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(key))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// @router /:uploadKey [delte]
func DeleteUpload(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.RemoveUpload(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON("Completed"))

	if err != nil {
		log.Println(err)
		return
	}
}
