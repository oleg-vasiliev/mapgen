package world

import (
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"mapgen/internal/world/grid"
)

const colorBytesSize = 4 // RGBA
const mapCellSizePx = 16 // 16px - real cell size, to use 16x16 tiles later

// World represents current world state including generated grid and it's config
type World struct {
	config Config
	grid   *grid.Grid

	gridCanvasBuff []byte        // bytes buffer for grid preview image
	gridCanvas     *ebiten.Image // grid preview in 1px per 1 cell scale
	worldCanvas    *ebiten.Image // normal size world image

	recreateRequired bool
}

func (w *World) Canvas() *ebiten.Image {
	return w.worldCanvas
}

func New(cfg Config) *World {
	// Default world settings
	w := &World{
		config: cfg,
	}

	w.gridCanvasBuff = make([]byte, w.config.Height*w.config.Width*colorBytesSize)
	// empty grid for world contents according to world size
	w.grid = grid.New(w.config.Width, w.config.Height)
	// gridCanvas where 1 cell corresponds to 1 colored pixel
	w.gridCanvas = ebiten.NewImage(w.config.Width, w.config.Height)
	// worldCanvas for drawing world state
	w.worldCanvas = ebiten.NewImage(w.config.Width*mapCellSizePx, w.config.Height*mapCellSizePx)
	w.recreateRequired = true
	return w
}

func (w *World) RecreateWorld() {
	heightsMap := newHeightsMap(w.config)
	w.grid.UpdateHeights(heightsMap)
	w.updateGridPreview()
	w.renderWorld()
}

func (w *World) updateGridPreview() {
	for id, cell := range w.grid.Cells {
		w.gridCanvasBuff[colorBytesSize*id+0] = cell.Color.R
		w.gridCanvasBuff[colorBytesSize*id+1] = cell.Color.G
		w.gridCanvasBuff[colorBytesSize*id+2] = cell.Color.B
		w.gridCanvasBuff[colorBytesSize*id+3] = cell.Color.A
	}
	// much faster than w.gridCanvas.Set(cell.x, cell.y, cell.Color)
	w.gridCanvas.WritePixels(w.gridCanvasBuff)
}

func (w *World) renderWorld() {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Scale(mapCellSizePx, mapCellSizePx)
	w.worldCanvas.Clear()
	w.worldCanvas.DrawImage(w.gridCanvas, &options)
	// draw additional things to world canvas here
	// TODO: Draw world objects - trees, houses, etc.
}

func (w *World) Update() {
	if w.recreateRequired {
		w.RecreateWorld()
		w.recreateRequired = false
	}
	w.listenControlUpdates()
}

func (w *World) listenControlUpdates() {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		w.config.RandSeed = rand.Int63()
		w.recreateRequired = true
	}

	// TODO: Add tooltips for noise settings controls

	// Noise Persistence
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		if w.config.Persistence > 0.1 {
			w.config.Persistence -= 0.1
			w.recreateRequired = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		w.config.Persistence += 0.1
		w.recreateRequired = true
	}
	// Noise Lacunarity
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		if w.config.Lacunarity > 0.01 {
			w.config.Lacunarity -= 0.01
			w.recreateRequired = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		w.config.Lacunarity += 0.01
		w.recreateRequired = true
	}

	// Noise Octaves
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		if w.config.Octaves > 1 {
			w.config.Octaves--
			w.recreateRequired = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		w.config.Octaves++
		w.recreateRequired = true
	}

	// Noise Scale
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		if w.config.Scale > 0.5 {
			w.config.Scale -= 0.5
			w.recreateRequired = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		w.config.Scale += 0.5
		w.recreateRequired = true
	}

}

func (w *World) NoiseScale() float64 {
	return w.config.Scale
}

func (w *World) NoisePersistence() float64 {
	return w.config.Persistence
}

func (w *World) NoiseLacunarity() float64 {
	return w.config.Lacunarity
}

func (w *World) NoiseOctaves() int {
	return w.config.Octaves
}

func (w *World) RandSeed() string {
	return strconv.FormatInt(w.config.RandSeed, 26)
}

func (w *World) WorldSize() (int, int) {
	return w.config.Width, w.config.Height
}

func (w *World) ConfigSeed() string {
	return w.config.StringSeed()
}

func (w *World) CanvasSnapshot() ebiten.Image {
	return *w.worldCanvas
}
