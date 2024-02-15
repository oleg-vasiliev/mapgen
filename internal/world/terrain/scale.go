package terrain

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"image/color"
	"io"
	"log"
	"strconv"
)

//go:embed sands.csv
var levels []byte

var lowest *Terrain

// TODO: Replace init with controllable flow
func init() {
	r := csv.NewReader(bytes.NewReader(levels))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		alt, err := strconv.ParseUint(record[0], 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		addTerrain(&Terrain{
			MaxAltitude: uint8(alt),
			Name:        record[2],
			Color:       parseHexColor(record[1]),
		})
	}
}

func addTerrain(new *Terrain) {
	// lowest is empty - new item becomes the first and lowest
	if lowest == nil {
		lowest = new
		return
	}
	// existing lowest becomes current and start iterating
	var cur, prev *Terrain
	cur = lowest
	// iterating
	for cur.next != nil {
		// new lower than current - place it before
		if new.MaxAltitude < cur.MaxAltitude {
			if prev == nil {
				new.next = cur
				lowest = new
				return
			} else {
				prev.next = new
				new.next = cur
				return
			}
		}
		// new larger (or equals to) than current -
		prev = cur
		cur = cur.next
	}
	// at this point all item have been checked - just add new item at the end
	cur.next = new
}

func FromAltitude(alt uint8) Terrain {
	var curr = lowest
	for {
		if alt <= curr.MaxAltitude {
			return *curr
		}
		if curr.next == nil {
			return *curr
		}
		if alt > curr.MaxAltitude {
			curr = curr.next
		}
	}
}

type Terrain struct {
	MaxAltitude uint8
	Name        string
	Color       color.RGBA
	next        *Terrain
}

func parseHexColor(hex string) color.RGBA {
	values, _ := strconv.ParseUint(hex[1:], 16, 32) // skipping leading #
	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}
}
