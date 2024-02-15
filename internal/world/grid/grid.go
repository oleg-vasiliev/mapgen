package grid

type Grid struct {
	Width  int
	Height int
	Cells  []Cell
}

func New(w, h int) *Grid {
	return &Grid{Width: w, Height: h, Cells: make([]Cell, w*h)}
}

// UpdateHeights fills grid Cells with heights map data to generate landscape
func (g *Grid) UpdateHeights(heightsMap []float64) {
	for id, height := range heightsMap {
		g.Cells[id] = NewCell(id%g.Width, id/g.Width, height)
	}
}
