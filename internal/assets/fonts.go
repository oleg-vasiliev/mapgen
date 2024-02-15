package assets

import (
	_ "embed"
	"log"

	"golang.org/x/image/font/opentype"
)

//go:embed ..\..\assets\fonts\PressStart2P-Regular.ttf
var fontSource []byte

var mainFont *opentype.Font

// TODO: Replace init with controlled loading
func init() {
	var err error
	mainFont, err = opentype.Parse(fontSource)
	if err != nil {
		log.Fatal("can't load font", err)
	}
}

func MainFont() *opentype.Font {
	return mainFont
}
