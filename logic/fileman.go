package logic

import (
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

func SaveFile(b io.Reader, header *multipart.FileHeader, info InfoHead) (key husk.Key, err error) {
	blob, mime, err := core.NewBLOB(b, info.For)

	if err != nil {
		return husk.CrazyKey(), err
	}

	upload := core.Upload{
		BLOB:     blob,
		Size:     header.Size,
		Name:     header.Filename,
		ItemKey:  info.ItemKey,
		ItemName: info.ItemName,
		MimeType: mime,
	}

	rec, err := upload.Create()

	if err != nil {
		return husk.CrazyKey(), err
	}

	return rec.GetKey(), nil
}
