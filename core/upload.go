package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"io"
	"mime/multipart"
)

type Upload struct {
	ItemKey  keys.TimeKey
	ItemName string `hsk:"size(75)"`
	Name     string `hsk:"size(50)"`
	MimeType string `hsk:"size(30)"`
	Size     int64
	BLOB     []byte `json:"-"` //Blob shouldn't be returned in JSON result sets.
}

func (u Upload) Valid() error {
	return validation.Struct(u)
}

func GetUploads(page, pagesize int) (records.Page, error) {
	return ctx.Uploads.Find(page, pagesize, op.Everything())
}

func GetUpload(key hsk.Key) (hsk.Record, error) {
	return ctx.Uploads.FindByKey(key)
}

func GetUploadFile(key hsk.Key) (result []byte, filename string, err error) {
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
	ItemKey  keys.TimeKey
	ItemName string
}

func GetInfoHead(header string) (InfoHead, error) {
	var result InfoHead
	err := json.Unmarshal([]byte(header), &result)

	return result, err
}

func SaveFile(b io.Reader, header *multipart.FileHeader, info InfoHead) (hsk.Key, error) {
	blob, mime, err := NewBLOB(b, info.For)

	if err != nil {
		return nil, err
	}

	upload := Upload{
		BLOB:     blob,
		Size:     header.Size,
		Name:     header.Filename,
		ItemKey:  info.ItemKey,
		ItemName: info.ItemName,
		MimeType: mime,
	}

	return upload.Create()
}

//GetUploadsBySize returns the first 50 records larger than @size bytes.
func GetUploadsBySize(size int64) (records.Page, error) {
	return ctx.Uploads.Find(1, 50, bySize(size))
}

func RemoveUpload(key hsk.Key) error {
	return ctx.Uploads.Delete(key)
}

func (u Upload) Create() (hsk.Key, error) {
	return ctx.Uploads.Create(u)
}
