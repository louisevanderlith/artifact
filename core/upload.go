package core

import (
	"github.com/louisevanderlith/husk"
)

type Upload struct {
	ItemKey  husk.Key
	ItemName string `hsk:"size(75)"`
	Name     string `hsk:"size(50)"`
	MimeType string `hsk:"size(30)"`
	Size     int64
	BLOB     []byte `json:"-"` //Blob shouldn't be returned in JSON result sets.
}

func (o Upload) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func GetUploads(page, pagesize int) husk.Collection {
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

	uploadData := upload.Data().(*Upload)
	blob := uploadData.BLOB

	return blob, uploadData.Name, err
}

//GetUploadsBySize returns the first 50 records larger than @size bytes.
func GetUploadsBySize(size int64) husk.Collection {
	return ctx.Uploads.Find(1, 50, bySize(size))
}

func RemoveUpload(key husk.Key) error {
	err := ctx.Uploads.Delete(key)

	if err != nil {
		return err
	}

	return ctx.Uploads.Save()
}

func (upload Upload) Create() (husk.Recorder, error) {
	rec := ctx.Uploads.Create(upload)

	if rec.Error != nil {
		return nil, rec.Error
	}

	err := ctx.Uploads.Save()

	if err != nil {
		return nil, err
	}

	return rec.Record, nil
}
