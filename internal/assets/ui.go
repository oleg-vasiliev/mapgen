package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
)

//go:embed ..\..\assets\ui\crosshair-square-b.png
var crosshairBytes []byte
var crosshairImage image.Image

//go:embed ..\..\assets\icon128.png
var mapIconBytes []byte
var mapIconBytesImage image.Image

func init() {
	var err error
	
	if crosshairImage, _, err = image.Decode(bytes.NewReader(crosshairBytes)); err != nil {
		log.Fatal("can't load crosshair image", err)
	}
	
	if mapIconBytesImage, _, err = image.Decode(bytes.NewReader(mapIconBytes)); err != nil {
		log.Fatal("can't load crosshair image", err)
	}
}

func CrosshairImage() image.Image {
	return crosshairImage
}

func MapIconImage() image.Image {
	return mapIconBytesImage
}
