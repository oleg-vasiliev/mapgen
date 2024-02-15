package main

import (
	"flag"
	"image"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"mapgen/internal/assets"
	"mapgen/internal/game"
	"mapgen/internal/overlay"
	"mapgen/internal/viewport"
	"mapgen/internal/world"
)

func main() {
	var (
		fullscreen bool
		configSeed string
		windowSize string
	)
	flag.BoolVar(&fullscreen, "fullscreen", false, "'--fullscreen' or '-f' for native fullscreen")
	flag.BoolVar(&fullscreen, "f", false, "'--fullscreen' or '-f' for native fullscreen")
	flag.StringVar(&windowSize, "window", "800x600", "--window 800x600 to set window resolution")
	flag.StringVar(&configSeed, "seed", "", "--seed 200x150@ghcbgo4ma1nmj")
	flag.Parse()

	var windowWidth, windowHeight int
	var err error
	{
		windowSizeSplit := strings.Split(windowSize, "x")
		if len(windowSizeSplit) != 2 {
			log.Fatal("invalid window size", err)
		}
		if windowWidth, err = strconv.Atoi(windowSizeSplit[0]); err != nil {
			log.Fatal("invalid window size", err)
		}
		if windowHeight, err = strconv.Atoi(windowSizeSplit[1]); err != nil {
			log.Fatal("invalid window size", err)
		}
	}

	if fullscreen {
		ebiten.SetFullscreen(true)
	} else {
		ebiten.SetWindowSize(windowWidth, windowHeight)
	}
	ebiten.SetWindowTitle("Mapgen")
	ebiten.SetWindowIcon([]image.Image{assets.MapIconImage()})

	config, err := world.NewConfig(world.WithSize(160, 90))
	if err != nil {
		log.Fatal("can't create configuration", err)
	}

	if len(configSeed) > 0 {
		err := config.ApplySeed(configSeed)
		if err != nil {
			log.Fatal("can't apply provided seed", err)
		}
	}

	w := world.New(config)
	o := overlay.New(w)
	v := viewport.New(w)
	g := game.New(w, o, v)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
