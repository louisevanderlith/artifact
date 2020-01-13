package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
	"github.com/louisevanderlith/artifact/logic"
)

func Get(c *gin.Context) {
	results := core.GetUploads(1, 10)

	c.JSON(http.StatusOK, results)
}

// @Title GetUploads
// @Description Gets the uploads
// @Success 200 {[]core.Upload} []core.Upload
// @router /all/:pagesize [get]
func Search(c *gin.Context) {
	ps := c.Param("pagesize")
	page, size := getPageData(ps)

	results := core.GetUploads(page, size)

	c.JSON(http.StatusOK, results)
}

// @Title GetUpload
// @Description Gets the requested upload
// @Param	uploadKey			path	husk.Key 	true		"Key of the file you require"
// @Success 200 {core.Upload} core.Upload
// @router /:uploadKey [get]
func View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	record, err := core.GetUpload(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, record)
}

// @Title UploadFile
// @Description Upload a file
// @Param    file        form     file    true        "File"
// @Param	body		body 	core.Upload	true		"body for upload content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func Create(c *gin.Context) {
	file, header, err := c.File("file")

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	defer file.Close()

	info := c.FindFormValue("info")
	infoHead, err := logic.GetInfoHead(info)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	log.Printf("Size: %v\tKey: %s\r", header.Size, infoHead.ItemKey.String())

	key, err := logic.SaveFile(file, header, infoHead)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, key)
}

func Update(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, nil)
}

// @router /:uploadKey [delte]
func Delete(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = core.RemoveUpload(key)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "Completed")
}

func getPageData(pageData string) (int, int) {
	defaultPage := 1
	defaultSize := 10

	if len(pageData) < 2 {
		return defaultPage, defaultSize
	}

	pChar := []rune(pageData[:1])

	if len(pChar) != 1 {
		return defaultPage, defaultSize
	}

	page := int(pChar[0]) % 32
	pageSize, err := strconv.Atoi(pageData[1:])

	if err != nil {
		return defaultPage, defaultSize
	}

	return page, pageSize
}