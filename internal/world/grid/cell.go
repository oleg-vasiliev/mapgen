package grid

import (
	"image/color"

	"mapgen/internal/world/terrain"
)

type Cell struct {
	X, Y        int   // cell coords in the grid
	Altitude    uint8 // 0-255
	Color       color.RGBA
	TerrainName string
}

func NewCell(x, y int, altitude float64) Cell {
	cell := Cell{X: x, Y: y}
	// Altitude already should be normalized for 0-1 range, but...
	if altitude < 0 {
		cell.Altitude = 0
	}
	if altitude > 1 {
		cell.Altitude = 255
	}
	cell.Altitude = byte(255 * altitude)
	// Calculate cell terrain type and color based on cell altitude
	// using some predefined adjustable set of rules for each terrain
	t := terrain.FromAltitude(cell.Altitude)
	cell.Color = t.Color
	cell.TerrainName = t.Name
	// for grayscale testing
	// cell.Color = color.RGBA{R: cell.Altitude, G: cell.Altitude, B: cell.Altitude, A: 255}
	// cell.TerrainName = "Grayscale"
	return cell
}
