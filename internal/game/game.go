package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"mapgen/internal/exporter"
	"mapgen/internal/overlay"
	"mapgen/internal/viewport"
	"mapgen/internal/world"
)

type Game struct {
	world    *world.World
	overlay  *overlay.Overlay
	viewport *viewport.Viewport
	exporter *exporter.ImageExporter
}

func New(w *world.World, o *overlay.Overlay, v *viewport.Viewport) *Game {
	return &Game{
		world:    w,
		overlay:  o,
		viewport: v,
		exporter: exporter.NewImageExporter(),
	}
}

func (g *Game) Layout(outerWidth, outerHeight int) (w, h int) {
	// optional logic to fix window size after window resizing to mod(cell_size)
	// dX := outerWidth % 16
	// dY := outerWidth % 16
	// if dX != 0 || dY!=0{
	// 	ebiten.SetWindowSize(outerWidth-dX, outerHeight-dY)
	// }

	// The unit of outsideWidth/Altitude is device-independent pixels.
	// By multiplying them by the device scale factor, we can get a hi-DPI screen size.
	// s := ebiten.DeviceScaleFactor()
	// return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)

	// logical screen size the same as physical
	g.viewport.Layout(outerWidth, outerHeight)
	g.overlay.Layout(outerWidth, outerHeight)
	return outerWidth, outerHeight
}

func (g *Game) Update() error {
	g.viewport.Update()
	g.world.Update()
	g.overlay.Update()
	// handle exit
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	// handle world image export
	// TODO: Add visual notification for image exporting
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		img := g.world.CanvasSnapshot()
		err := g.exporter.Export(&img, g.world.ConfigSeed())
		if err != nil {
			log.Print("error exporting world image", err)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 80})
	g.viewport.Draw(screen)
	g.overlay.Draw(screen)
}
