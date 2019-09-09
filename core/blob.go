package core

import (
	"bufio"
	"bytes"
	"errors"
	"image"

	"github.com/disintegration/imaging"
	"github.com/louisevanderlith/artifact/core/optimizetype"
)

type optimFunc func(data image.Image) (result []byte, mimetype string, err error)
type optmizer map[optimizetype.Enum]optimFunc

var optimizers optmizer

func init() {
	optimizers = getOptimizers()
}

func NewBLOB(data []byte, purpose string) ([]byte, string, error) {
	targetType := optimizetype.GetEnum(purpose)
	return OptimizeFor(data, targetType)
}

//OptimizeFor returns the new Bytes, MIME Type and an error
func OptimizeFor(data []byte, oType optimizetype.Enum) ([]byte, string, error) {
	reader := bytes.NewReader(data)
	decoded, err := imaging.Decode(reader)

	if err != nil {
		return nil, "", err
	}

	opt, hasOpt := optimizers[oType]

	if !hasOpt {
		return nil, "", errors.New("optimizer Type not found")
	}

	return opt(decoded)
}

func getOptimizers() optmizer {
	result := make(optmizer)

	result[optimizetype.Logo] = optimizeLogo
	result[optimizetype.Banner] = optimizeBanner
	result[optimizetype.Ad] = optimizeAd
	result[optimizetype.Thumb] = optimizeThumb

	return result
}

func optimizeAd(data image.Image) ([]byte, string, error) {
	return optimize(data, 700, 450, imaging.JPEG)
}

func optimizeBanner(data image.Image) ([]byte, string, error) {
	return optimize(data, 1536, 864, imaging.JPEG)
}

func optimizeLogo(data image.Image) ([]byte, string, error) {
	return optimize(data, 256, 128, imaging.PNG)
}

func optimizeThumb(data image.Image) ([]byte, string, error) {
	return optimize(data, 350, 145, imaging.JPEG)
}

func optimize(data image.Image, width, height int, format imaging.Format) ([]byte, string, error) {
	var b bytes.Buffer

	writer := bufio.NewWriter(&b)
	defer writer.Flush()

	optImage := imaging.Fit(data, width, height, imaging.Lanczos)

	err := imaging.Encode(writer, optImage, format)

	if err != nil {
		return nil, "", err
	}

	mimetype := "image/" + format.String()

	return b.Bytes(), mimetype, nil
}
