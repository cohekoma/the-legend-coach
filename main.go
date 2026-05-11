package main

import (
	"log"

	"github.com/cohekoma/the-legend-coach/internals/config"
	"github.com/cohekoma/the-legend-coach/internals/director"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.StartupWidth, config.StartupHeight)
	ebiten.SetWindowTitle(config.GameTitle)

	if err := ebiten.RunGame(director.CreateNewGame()); err != nil {
		log.Fatal("Error")
	}
}
