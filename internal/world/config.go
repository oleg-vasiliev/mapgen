package world

import (
	"fmt"
	"strconv"
	"strings"
)

const MaxWorldSize = 1000 // artificial limitation until the mini map and accelerated scrolling are implemented
const MinWorldSize = 10   // lesser size just does not make sense

// Config represents all world settings including noise permutation and seed
// TODO: Add scale/octaves/persistence/lacunarity values to seed string
type Config struct {
	Width    int
	Height   int
	RandSeed int64

	Scale       float64
	Octaves     int
	Persistence float64
	Lacunarity  float64
	Offset      vector // octave offset persistent part
}

func NewConfig(opts ...Option) (Config, error) {
	// Default configuration without any options
	c := Config{
		Width:    60,
		Height:   40,
		RandSeed: 1591016428052150547,

		Scale:       30.0,
		Octaves:     2,
		Persistence: 2.0,
		Lacunarity:  0.3,
		Offset:      vector{x: -100, y: 0.001},
	}

	// Apply config options
	for _, opt := range opts {
		opt(&c)
	}

	// Validate configuration
	if c.Width < MinWorldSize || c.Width > MaxWorldSize {
		return Config{}, fmt.Errorf("invalid world width: %d", c.Width)
	}
	if c.Height < MinWorldSize || c.Height > MaxWorldSize {
		return Config{}, fmt.Errorf("invalid world height: %d", c.Height)
	}
	return c, nil
}

type Option func(c *Config)

// possible implementation where options validate themselves
// type Option func(c *Config) error

func WithSize(w, h int) Option {
	return func(c *Config) {
		c.Width = w
		c.Height = h
	}
}
func WithRandSeed(seed int64) Option {
	return func(c *Config) {
		c.RandSeed = seed
	}
}

func (s *Config) StringSeed() string {
	return fmt.Sprintf("%dx%d@%s", s.Width, s.Height, strconv.FormatInt(s.RandSeed, 26))
}

func (s *Config) ApplySeed(str string) error {
	sizeToRnd := strings.Split(str, "@")
	if len(sizeToRnd) != 2 {
		return fmt.Errorf("invalid seed string: %s", str)
	}
	// Parse randomization seed
	seed, err := strconv.ParseInt(sizeToRnd[1], 26, 64)
	if err != nil {
		return fmt.Errorf("invalid seed string: %s", str)
	}
	s.RandSeed = seed

	// Parse dimensions
	dimensionsSplit := strings.Split(sizeToRnd[0], "x")
	if len(dimensionsSplit) != 2 {
		return fmt.Errorf("invalid seed dimensions substring: %s", sizeToRnd[0])
	}
	if s.Width, err = strconv.Atoi(dimensionsSplit[0]); err != nil {
		return fmt.Errorf("%w: invalid dimensions: %s", err, str)
	}
	if s.Height, err = strconv.Atoi(dimensionsSplit[1]); err != nil {
		return fmt.Errorf("%w: invalid dimensions: %s", err, str)
	}
	// Validate dimensions
	if s.Width < MinWorldSize || s.Width > MaxWorldSize {
		return fmt.Errorf("invalid world width size: %d", s.Width)
	}
	if s.Height < MinWorldSize || s.Height > MaxWorldSize {
		return fmt.Errorf("invalid world height size: %d", s.Height)
	}
	return nil
}
