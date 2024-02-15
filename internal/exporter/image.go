package exporter

import (
	"fmt"
	"image/jpeg"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageExporter struct {
	m sync.Mutex
}

func NewImageExporter() *ImageExporter {
	return &ImageExporter{}
}

func (ie *ImageExporter) Export(img *ebiten.Image, name string) error {
	if !ie.m.TryLock() {
		return fmt.Errorf("export already in progress")
	}
	defer ie.m.Unlock()
	file, err := os.Create(name + ".jpg")
	if err != nil {
		return err
	}
	defer file.Close()
	if err = jpeg.Encode(file, img, nil); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}
	return nil
}
