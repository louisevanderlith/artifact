package logic

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/artifact/core"
)

type InfoHead struct {
	For      string
	ItemKey  husk.Key
	ItemName string
}

func GetInfoHead(header string) (InfoHead, error) {
	var result InfoHead
	err := json.Unmarshal([]byte(header), &result)

	return result, err
}

func SaveFile(file multipart.File, header *multipart.FileHeader, info InfoHead) (key husk.Key, err error) {
	var b bytes.Buffer
	copied, err := io.Copy(&b, file)

	if err != nil {
		return husk.CrazyKey(), err
	}

	blob, mime, err := core.NewBLOB(b.Bytes(), info.For)

	if err != nil {
		return husk.CrazyKey(), err
	}

	upload := core.Upload{
		BLOB:     blob,
		Size:     copied,
		Name:     header.Filename,
		ItemKey:  info.ItemKey,
		ItemName: info.ItemName,
		MimeType: mime,
	}

	rec := upload.Create()

	if rec.Error != nil {
		return husk.CrazyKey(), rec.Error
	}

	return rec.Record.GetKey(), nil
}
