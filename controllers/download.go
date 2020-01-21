package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/husk"
)

// @Title GetFile
// @Description Gets the requested file only
// @Param	uploadID			path	int64 	true		"ID of the file you require"
// @Success 200 {[]byte} []byte
// @router /file/:uploadKey [get]
func Download(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result, fileName, err := core.GetUploadFile(key)

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusNotFound, nil)
		return
	}

	ext := getExt(fileName)
	mimes := make(map[string]string)
	mimes["js"] = "text/javascript"
	mimes["css"] = "text/css"
	mimes["html"] = "text/html"
	mimes["ico"] = "image/x-icon"
	mimes["font"] = "font/" + ext
	mimes["jpeg"] = "image/jpeg"
	mimes["jpg"] = "image/jpeg"
	mimes["png"] = "image/png"

	content := mimes[ext]

	c.DataFromReader(http.StatusOK, int64(len(result)), content, bytes.NewReader(result), headers(content, fileName))
}

func headers(contenttype, filename string) map[string]string {
	result := make(map[string]string)

	result["Strict-Transport-Security"] = "max-age=31536000; includeSubDomains"
	result["Access-Control-Allow-Credentials"] = "true"
	result["Server"] = "kettle"
	result["X-Content-Type-Options"] = "nosniff"

	result["Content-Description"] = "File Transfer"
	result["Content-Transfer-Encoding"] = "binary"
	result["Expires"] = "0"
	result["Cache-Control"] = "must-revalidate"
	result["Pragma"] = "public"

	result["Content-Disposition"] = "attachment; filename=" + filename
	result["Content-Type"] = contenttype

	return result
}

func getExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	return filename[dotIndex+1:]
}
