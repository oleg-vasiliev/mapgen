package overlay

type WorldStatsProvider interface {
	NoiseScale() float64
	NoisePersistence() float64
	NoiseLacunarity() float64
	NoiseOctaves() int

	RandSeed() string
	WorldSize() (int, int)
}
