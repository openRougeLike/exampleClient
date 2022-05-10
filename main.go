package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return DIMENSIONS * CurrentResourcePack.TileSize, DIMENSIONS * CurrentResourcePack.TileSize
}

func main() {
	m := Init()
	game := Game{
		Map: m,
	}

	LoadResourcePack()

	ebiten.SetWindowSize(DIMENSIONS*CurrentResourcePack.TileSize, DIMENSIONS*CurrentResourcePack.TileSize)
	ebiten.SetWindowTitle("Poggers Game")

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
