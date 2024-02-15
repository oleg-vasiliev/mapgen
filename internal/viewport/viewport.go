package viewport

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const scrollSpeed = 20
const zoomSpeed = 0.05
const maxZoom = 2
const minZoom = 0.5

// Viewport tracks own position and know how to draw specified area
// of the world to the screen, so world object should be injected
type Viewport struct {
	world         WorldRenderer
	width, height int     // actual viewport size in pixels
	posX, posY    int     // viewport position on the map (top-left corner based coords)
	scale         float64 // view scale factor
	// velX, velY    int // velocity for smooth movement implementation
}

func New(w WorldRenderer) *Viewport {
	return &Viewport{
		world:  w,
		width:  0,
		height: 0,
		scale:  1,
		posX:   0,
		posY:   0,
	}
}

type WorldRenderer interface {
	Canvas() *ebiten.Image
}

func (v *Viewport) Layout(availableWidth, availableHeight int) {
	v.width = availableWidth
	v.height = availableHeight
}

func (v *Viewport) Draw(screen *ebiten.Image) {
	// Draw world image to viewport according to viewport settings
	options := ebiten.DrawImageOptions{}
	// apply position
	options.GeoM.Translate(float64(v.posX), float64(v.posY))
	// apply scaling - move map half-screen up-left to center origin point
	// scale it, then move back to original position
	options.GeoM.Translate(-float64(v.width/2), -float64(v.height/2))
	options.GeoM.Scale(v.scale, v.scale)
	options.GeoM.Translate(float64(v.width/2), float64(v.height/2))
	// Actually draw a world
	screen.DrawImage(v.world.Canvas(), &options)
}

func (v *Viewport) Update() {
	// Simplest viewport movement implementation
	// TODO: Improvement - implement inertial scroll with acceleration
	if inpututil.KeyPressDuration(ebiten.KeyQ) > 0 {
		if v.scale > minZoom {
			v.scale -= zoomSpeed
		}
	}
	if inpututil.KeyPressDuration(ebiten.KeyE) > 0 {
		if v.scale < maxZoom {
			v.scale += zoomSpeed
		}
	}
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		v.posY += scrollSpeed
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		v.posY -= scrollSpeed
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		v.posX += scrollSpeed
	}
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		v.posX -= scrollSpeed
	}
}
