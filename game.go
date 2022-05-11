package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Screen int8

const (
	S_Tile Screen = iota
	S_Lore
)

type Game struct {
	Map Map
	Direction Direction
	Screen Screen
	NPC NPCProg
}

type NPCProg struct {
	Text []string
	TickSinceStart int
}

func (g *Game) Update() error {
	var err *APIError

	if keyHeld(ebiten.KeyR) {
		m := Init()
		g.Map = m
		g.Screen = S_Tile
		return nil
	}

	switch g.Screen {
	case S_Tile:
		dirs := map[Direction]bool{}

		if keyHeld(ebiten.KeyZ) {
			err = g.ZKey()
		}
	
		dirs[DirUp] = keyHeld(ebiten.KeyArrowUp)
		dirs[DirRight] = keyHeld(ebiten.KeyArrowRight)
		dirs[DirDown] = keyHeld(ebiten.KeyArrowDown)
		dirs[DirLeft] = keyHeld(ebiten.KeyArrowLeft)
	
		err = g.ArrowKey(dirs)
	
		if err != nil {
			fmt.Println(err)
		}
	
		if g.Screen == S_Lore {
			g.NPC.TickSinceStart++
		}
	case S_Lore:
		
	}



	return nil
}

func keyHeld(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d != 0 && d%5 == 0
}

func (g *Game) SetScreen(newScreen Screen) {
	g.Screen = newScreen
	fmt.Println("New screen!")
}