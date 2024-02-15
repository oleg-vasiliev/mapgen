package overlay

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"mapgen/internal/assets"
)

// TODO: Display current terrain levels configuration from world/terrain pkg

const dpi = 72

var panelsColor = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xAA}
var textColor = color.Gray{Y: 220}

func New(worldStats WorldStatsProvider) *Overlay {
	overlayFontFace, err := opentype.NewFace(
		assets.MainFont(),
		&opentype.FaceOptions{
			Size:    12,
			DPI:     dpi,
			Hinting: font.HintingVertical,
		})
	if err != nil {
		log.Fatal(err)
	}

	faceWithLineHeight := text.FaceWithLineHeight(overlayFontFace, 16)
	o := Overlay{
		wordsStats:        worldStats,
		fontFace:          faceWithLineHeight,
		overlayIsHidden:   false,
		controlsHintPanel: ebiten.NewImage(280, 110),
		worldDetailsPanel: ebiten.NewImage(280, 160),
		hiddenOverlayHint: ebiten.NewImage(280, 26),
		crosshair:         ebiten.NewImageFromImage(assets.CrosshairImage()),
	}
	return &o
}

type Overlay struct {
	wordsStats      WorldStatsProvider
	fontFace        font.Face
	overlayIsHidden bool

	controlsHintPanel *ebiten.Image
	worldDetailsPanel *ebiten.Image
	hiddenOverlayHint *ebiten.Image
	crosshair         *ebiten.Image

	screenWidth  int
	screenHeight int
}

func (o *Overlay) Draw(screen *ebiten.Image) {
	if !o.overlayIsHidden {
		o.drawControlsHintPanel(screen)
		o.drawWorldDetailsPanel(screen)
	} else {
		o.drawHiddenOverlayHintPanel(screen)
	}
	o.drawCrosshair(screen)
}

func (o *Overlay) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		o.overlayIsHidden = !o.overlayIsHidden
	}
}

func (o *Overlay) Layout(availableWidth, availableHeight int) {
	o.screenWidth = availableWidth
	o.screenHeight = availableHeight
}

func (o *Overlay) drawControlsHintPanel(screen *ebiten.Image) {
	o.controlsHintPanel.Fill(panelsColor)
	text.Draw(o.controlsHintPanel, overlayControlsTemplate, o.fontFace, 10, 20, textColor)
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, float64(o.screenHeight-o.controlsHintPanel.Bounds().Dy()))
	screen.DrawImage(o.controlsHintPanel, &options)
}

func (o *Overlay) drawHiddenOverlayHintPanel(screen *ebiten.Image) {
	o.hiddenOverlayHint.Fill(panelsColor)
	text.Draw(o.hiddenOverlayHint, "(!) [Tab] to show UI", o.fontFace, 10, 20, textColor)
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, float64(o.screenHeight-o.hiddenOverlayHint.Bounds().Dy()))
	screen.DrawImage(o.hiddenOverlayHint, &options)
}

func (o *Overlay) drawWorldDetailsPanel(screen *ebiten.Image) {
	o.worldDetailsPanel.Fill(panelsColor)
	worldWidth, worldHeight := o.wordsStats.WorldSize()
	t := fmt.Sprintf(overlayDetailsTemplate,
		worldWidth, worldHeight, o.wordsStats.RandSeed(),
		o.wordsStats.NoisePersistence(), o.wordsStats.NoiseLacunarity(), o.wordsStats.NoiseOctaves(), o.wordsStats.NoiseScale())
	text.Draw(o.worldDetailsPanel, t, o.fontFace, 10, 20, textColor)
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(o.screenWidth-o.worldDetailsPanel.Bounds().Dx()), 0)
	screen.DrawImage(o.worldDetailsPanel, &options)
}

func (o *Overlay) drawCrosshair(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64((o.screenWidth-o.crosshair.Bounds().Dx())/2),
		float64((o.screenHeight-o.crosshair.Bounds().Dy())/2))
	screen.DrawImage(o.crosshair, &options)
}

const overlayControlsTemplate = `Controls help:
·[WASD] scroll view
·[E/Q] zoom in/out
·[Tab] toggle UI
·[Space] random world
·[Esc] exit
`

const overlayDetailsTemplate = `World details:
·Dimensions: %dx%d
·Seed: %s

Noise settings:
·Persistence...%0.2f
·Lacunarity....%0.2f
·Octaves.......%d
·Scale.........%0.2f
`
