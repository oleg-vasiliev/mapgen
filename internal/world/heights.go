package world

import (
	"math"
	"math/rand"

	"mapgen/pkg/simplex"
)

type vector struct {
	x, y float64
}

func newHeightsMap(cfg Config) []float64 {
	random := rand.New(rand.NewSource(cfg.RandSeed))
	noiseMap := make([]float64, cfg.Width*cfg.Height)

	// Fix possible invalid Config
	if cfg.Octaves < 1 {
		cfg.Octaves = 1
	}
	if cfg.Scale <= 0 {
		cfg.Scale = 0.0001
	}

	octaveOffsets := make([]*vector, 0)
	for i := 1; i <= cfg.Octaves; i++ {
		octaveOffsets = append(octaveOffsets, &vector{
			x: randFloat(random, -100000, 100000) + cfg.Offset.x,
			y: randFloat(random, -100000, 100000) + cfg.Offset.y,
		})
	}

	maxNoiseHeight := math.SmallestNonzeroFloat64
	minNoiseHeight := math.MaxFloat64

	// When changing noise scale, it zooms from top-right corner
	// This will make it zoom from the center
	halfWidth := float64(cfg.Width / 2)
	halfHeight := float64(cfg.Height / 2)

	for x := 0; x < cfg.Width; x++ {
		for y := 0; y < cfg.Height; y++ {
			amplitude := float64(1)
			frequency := float64(1)
			noiseHeight := float64(0)

			// Calculate noise for each octave
			for _, offset := range octaveOffsets {
				// We sample a point (x,y)
				sampleX := (float64(x)-halfWidth)/cfg.Scale*frequency + offset.x
				sampleY := (float64(y)-halfHeight)/cfg.Scale*frequency + offset.y

				noiseValue := simplex.Noise2(sampleX, sampleY)*2 - 1

				// noiseHeight is our final noise, we add all octaves together here
				noiseHeight += noiseValue * amplitude
				amplitude *= cfg.Persistence
				frequency *= cfg.Lacunarity
			}
			// Find the min and max noise height in noise map
			// to interpolate the min and max values between 0 and 1 again
			if noiseHeight > maxNoiseHeight {
				maxNoiseHeight = noiseHeight
			} else if noiseHeight < minNoiseHeight {
				minNoiseHeight = noiseHeight
			}
			// Assign noise
			noiseMap[y*cfg.Width+x] = noiseHeight
		}
	}

	for x := 0; x < cfg.Width; x++ {
		for y := 0; y < cfg.Height; y++ {
			// Returns a value between 0f and 1f based on noiseMap value
			// minNoiseHeight being 0f, and maxNoiseHeight being 1f
			noiseMap[y*cfg.Width+x] = inverseLerp(minNoiseHeight, maxNoiseHeight, noiseMap[y*cfg.Width+x])
		}
	}
	return noiseMap
}

func randFloat(random *rand.Rand, min, max float64) float64 {
	return min + random.Float64()*(max-min)
}

func inverseLerp(a0, a1, w float64) float64 {
	return (w - a0) / (a1 - a0)
}

func lerp(a0, a1, w float64) float64 {
	// if 0.0 > width {
	//	return a0
	// }
	// if 1.0 < width {
	//	return a1
	// }

	// Use this cubic interpolation [[Smoothstep]] instead, for a smooth appearance:
	// return (a1-a0)*(3.0-width*2.0)*width*width + a0

	// Use [[Smootherstep]] for an even smoother result with a second derivative equal to zero on boundaries:
	// return (a1 - a0) * ((width * (width * 6.0 - 15.0) + 10.0) * width * width * width) + a0;

	return (a1-a0)*w + a0
}
