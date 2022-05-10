package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Map Map
}

func (g *Game) Update() error {
	var err *APIError

	dirs := map[Direction]bool{}

	dirs[DirUp] = keyHeld(ebiten.KeyArrowUp)
	dirs[DirRight] = keyHeld(ebiten.KeyArrowRight)
	dirs[DirDown] = keyHeld(ebiten.KeyArrowDown)
	dirs[DirLeft] = keyHeld(ebiten.KeyArrowLeft)

	err = g.ArrowKey(dirs)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func keyHeld(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d != 0 && d%5 == 0
}
