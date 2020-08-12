package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
)

func GetUploads(w http.ResponseWriter, r *http.Request) {
	results, err := core.GetUploads(1, 10)

	if err != nil {
		log.Println("GetUploads Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(results))

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
	page, size := drx.GetPageData(r)

	results, err := core.GetUploads(page, size)

	if err != nil {
		log.Println("GetUploads Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// ViewUpload
func ViewUpload(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

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

	err = mix.Write(w, mix.JSON(record))

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
	file, header, err := r.FormFile("file")

	if err != nil {
		log.Println("File Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer file.Close()

	info := r.FormValue("info")
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

	err = mix.Write(w, mix.JSON(key))

	if err != nil {
		log.Println("Serve Error", err)
		return
	}
}

// @router /:uploadKey [delte]
func DeleteUpload(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

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

	err = mix.Write(w, mix.JSON("Completed"))

	if err != nil {
		log.Println(err)
		return
	}
}
