package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"io"
	"mime/multipart"
)

type Upload struct {
	ItemKey  husk.Key
	ItemName string `hsk:"size(75)"`
	Name     string `hsk:"size(50)"`
	MimeType string `hsk:"size(30)"`
	Size     int64
	BLOB     []byte `json:"-"` //Blob shouldn't be returned in JSON result sets.
}

func (u Upload) Valid() (bool, error) {
	return husk.ValidateStruct(&u)
}

func GetUploads(page, pagesize int) (husk.Collection, error) {
	return ctx.Uploads.Find(page, pagesize, husk.Everything())
}

func GetUpload(key husk.Key) (husk.Recorder, error) {
	return ctx.Uploads.FindByKey(key)
}

func GetUploadFile(key husk.Key) (result []byte, filename string, err error) {
	upload, err := GetUpload(key)

	if err != nil {
		return nil, "", err
	}

	uploadData := upload.Data().(Upload)
	blob := uploadData.BLOB

	return blob, uploadData.Name, err
}

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
	blob, mime, err := NewBLOB(b, info.For)

	if err != nil {
		return husk.CrazyKey(), err
	}

	upload := Upload{
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

//GetUploadsBySize returns the first 50 records larger than @size bytes.
func GetUploadsBySize(size int64) (husk.Collection, error) {
	return ctx.Uploads.Find(1, 50, bySize(size))
}

func RemoveUpload(key husk.Key) error {
	err := ctx.Uploads.Delete(key)

	if err != nil {
		return err
	}

	return ctx.Uploads.Save()
}

func (u Upload) Create() (husk.Recorder, error) {
	rec := ctx.Uploads.Create(u)

	if rec.Error != nil {
		return nil, rec.Error
	}

	err := ctx.Uploads.Save()

	if err != nil {
		return nil, err
	}

	return rec.Record, nil
}
