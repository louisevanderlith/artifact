package core

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/louisevanderlith/artifact/core/optimizetype"
)

func getImage(location string) []byte {
	dat, err := ioutil.ReadFile(location)

	if err != nil {
		log.Println(err)
	}

	return dat
}

func saveImage(location string, file []byte) {
	ioutil.WriteFile(location, file, 0644)
}

func TestBlob_OptimizeFor_Logo_PNG2PNG(t *testing.T) {
	resultName := "png2png_logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.png")

	bits, _, err := OptimizeFor(data, optimizetype.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Logo_JPG2PNG(t *testing.T) {
	resultName := "jpg2png_logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.jpg")

	bits, _, err := OptimizeFor(data, optimizetype.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Banner_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.png")

	bits, _, err := OptimizeFor(data, optimizetype.Banner)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Banner_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_banner.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.jpg")

	bits, _, err := OptimizeFor(data, optimizetype.Banner)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Ad_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.png")

	bits, _, err := OptimizeFor(data, optimizetype.Ad)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Ad_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_ad.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.jpg")

	bits, _, err := OptimizeFor(data, optimizetype.Ad)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Thumb_PNG2JPG(t *testing.T) {
	resultName := "png2jpg_thumb.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.png")

	bits, _, err := OptimizeFor(data, optimizetype.Thumb)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Thumb_JPG2JPG(t *testing.T) {
	resultName := "jpg2jpg_thumb.jpg"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./test.jpg")

	bits, _, err := OptimizeFor(data, optimizetype.Thumb)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	saveImage(writeLocation, bits)
}

func TestBlob_OptimizeFor_Ad(t *testing.T) {
	resultName := "logo.png"
	writeLocation := "./testData/" + resultName
	os.Remove(writeLocation)

	data := getImage("./logo.png")

	bits, _, err := OptimizeFor(data, optimizetype.Logo)

	if err != nil {
		t.Error("Error occurred:", err)
	}

	t.Log(len(bits))
	if len(bits) > 0 {
		saveImage(writeLocation, bits)
	} else {
		t.Error("Image Zero")
	}

}
