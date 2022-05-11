package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileInfoGen struct {
	North,
	East,
	South,
	West,
	NE,
	NW,
	SE,
	SW bool
}

var TileMap = map[string][2]int{
	"1101_1111": {0, 0},
	"1000_1101": {0, 1},
	"1000_1011": {0, 2},
	"0001_0111": {0, 3},
	"0100_1110": {0, 4},
	"0000_0100": {0, 5},
	"0010_0110": {0, 6},
	"0000_0010": {0, 7},

	"0101_1111": {1, 0},
	"0010_1110": {1, 1},
	"0010_0111": {1, 2},
	"0001_1011": {1, 3},
	"0100_1101": {1, 4},
	"0100_1100": {1, 5},
	"1111_1111": {1, 6},
	"0001_0011": {1, 7},

	"0111_1111": {2, 0},
	"1011_1111": {2, 1},
	"1010_1111": {2, 2},
	"1110_1111": {2, 3},
	"arena":	 {2, 4},
	"0000_1000": {2, 5},
	"1000_1001": {2, 6},
	"0000_0001": {2, 7},

	"1001_1111": {3, 0},
	"1000_1111": {3, 1},
	"1100_1111": {3, 2},
	"1001_1011": {3, 3},
	"0000_1001": {3, 4},
	"1100_1101": {3, 5},
	"0000_1110": {3, 6},
	"0000_0111": {3, 7},

	"0001_1111": {4, 0},
	"0000_1111": {4, 1},
	"0100_1111": {4, 2},
	"0000_0011": {4, 3},
	"0000_0000": {4, 4},
	"0000_1100": {4, 5},
	"0000_1101": {4, 6},
	"0000_1011": {4, 7},

	"0011_1111": {5, 0},
	"0010_1111": {5, 1},
	"0110_1111": {5, 2},
	"0011_0111": {5, 3},
	"0000_0110": {5, 4},
	"0110_1110": {5, 5},
	"0000_1010": {5, 6},
	"0000_0101": {5, 7},

	"chest":           {6, 0},
	"chest_open":      {6, 1},
	"chest_rare":      {6, 2},
	"chest_rare_open": {6, 3},
	"exit":            {6, 5},
	"monster":           {6, 6},
	"monster_done":      {6, 7},

	"npc_weapon":  {7, 0},
	"npc_armor":   {7, 1},
	"npc_potion":  {7, 2},
	"npc_upgrade": {7, 3},
	"npc_lore_0":  {7, 4},
	"npc_lore_1":  {7, 5},
	"npc_lore_2":  {7, 6},
	"npc_lore_3":  {7, 7},

	"npc_lore_4":   {8, 0},
	"npc_lore_5":   {8, 1},
	"npc_lore_6":   {8, 2},
	"npc_lore_7":   {8, 3},
	"npc_lore_8":   {8, 4},
	"npc_lore_9":   {8, 5},

	"user_0": 		{9,0},
	"user_1": 		{9,1},
	"user_2": 		{9,2},
	"user_3": 		{9,3},

	"SP_npcLore": 	{0, 0},
}

func BoolString(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func (t *TileInfoGen) String() (out string) {
	if t.North {
		t.NE = true
		t.NW = true
	}

	if t.South {
		t.SE = true
		t.SW = true
	}

	if t.East {
		t.NE = true
		t.SE = true
	}

	if t.West {
		t.NW = true
		t.SW = true
	}

	out += BoolString(t.North)
	out += BoolString(t.East)
	out += BoolString(t.South)
	out += BoolString(t.West)
	out += "_"
	out += BoolString(t.NE)
	out += BoolString(t.SE)
	out += BoolString(t.SW)
	out += BoolString(t.NW)

	return
}

func drawImg(screen *ebiten.Image, src *ebiten.Image, srcStartX, srcStartY, srcEndX, srcEndY, dstX, dstY int) {
	srcRect := image.Rect(srcStartX, srcStartY, srcEndX, srcEndY)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(dstX), float64(dstY))
	screen.DrawImage(src.SubImage(srcRect).(*ebiten.Image), op)
}

func DrawTile(screen *ebiten.Image, tileRef string, layoutX int, layoutY int) {
	if resourceLoc, ok := TileMap[tileRef]; ok {
		drawImg(screen, CurrentResourcePack.Tiles,
			resourceLoc[1]*CurrentResourcePack.TileSize,
			resourceLoc[0]*CurrentResourcePack.TileSize,
			(resourceLoc[1]+1)*CurrentResourcePack.TileSize, 
			(resourceLoc[0]+1)*CurrentResourcePack.TileSize,
			layoutX*CurrentResourcePack.TileSize,
			layoutY*CurrentResourcePack.TileSize,
		)
	} else {
		panic(tileRef)
	}
}

func DrawSprite(screen *ebiten.Image, sprite string, dstX, dstY int) {
	if resourceLoc, ok := TileMap[sprite]; ok {
		drawImg(screen, CurrentResourcePack.Sprites,
			resourceLoc[1]*CurrentResourcePack.SpriteSize,
			resourceLoc[0]*CurrentResourcePack.SpriteSize,
			(resourceLoc[1]+1)*CurrentResourcePack.SpriteSize, 
			(resourceLoc[0]+1)*CurrentResourcePack.SpriteSize,
			dstX,
			dstY,
		)
	} else {
		panic(sprite)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Screen {
	case S_Tile:
		DrawTileScreen(g, screen)
	case S_Lore:
		DrawDialogEntities([]string{"SP_npcLore"}, screen)
	}
	
}

func DrawTileScreen(g *Game, screen *ebiten.Image) {
	for layoutY := 0; layoutY < DIMENSIONS; layoutY++ {
		for layoutX := 0; layoutX < DIMENSIONS; layoutX++ {
			tileImg := "1111_1111"

			tile := g.Map.Layout[layoutY][layoutX]

			if tile != MapWall {
				strBuilder := TileInfoGen{}

				if layoutY != 0 {
					strBuilder.North = g.Map.Layout[layoutY-1][layoutX] == MapWall

					if layoutX != 0 {
						strBuilder.NW = g.Map.Layout[layoutY-1][layoutX-1] == MapWall
					}
					if layoutX != DIMENSIONS-1 {
						strBuilder.NE = g.Map.Layout[layoutY-1][layoutX+1] == MapWall
					}
				} else {
					strBuilder.North = true
				}

				if layoutY != DIMENSIONS-1 {
					strBuilder.South = g.Map.Layout[layoutY+1][layoutX] == MapWall
					if layoutX != 0 {
						strBuilder.SW = g.Map.Layout[layoutY+1][layoutX-1] == MapWall
					}
					if layoutX != DIMENSIONS-1 {
						strBuilder.SE = g.Map.Layout[layoutY+1][layoutX+1] == MapWall
					}
				} else {
					strBuilder.South = true
				}

				if layoutX != 0 {
					strBuilder.West = g.Map.Layout[layoutY][layoutX-1] == MapWall
				} else {
					strBuilder.West = true
				}

				if layoutX != DIMENSIONS-1 {
					strBuilder.East = g.Map.Layout[layoutY][layoutX+1] == MapWall
				} else {
					strBuilder.East = true
				}

				tileImg = strBuilder.String()
			}

			DrawTile(screen, tileImg, layoutX, layoutY)
		}
	}

	DrawTile(screen, "user_" + fmt.Sprint(g.Direction), g.Map.User[1], g.Map.User[0])
	DrawTile(screen, "exit", g.Map.Exit[1], g.Map.Exit[0])

	for str, info := range g.Map.ChestsInfo {
		coords := StringToCoords(str)
		tileRef := "chest"
		if info.Rare {
			tileRef += "_rare"
		}
		if info.Open {
			tileRef += "_open"
		}
		DrawTile(screen, tileRef, coords[1], coords[0])
	}

	for str, info := range g.Map.Rooms {
		coords := StringToCoords(str)
		tileRef := "arena"
		if info {
			tileRef += "_done"
		}

		DrawTile(screen, tileRef, coords[1], coords[0])
	}

	for str, t := range g.Map.NPC {
		coords := StringToCoords(str)
		tileRef := "npc_"

		switch t {
		case NPCShopWeapon:
			tileRef += "weapon"
		case NPCShopArmor:
			tileRef += "armor"
		case NPCShopPotion:
			tileRef += "potion"
		case NPCUpgrade:
			tileRef += "upgrade"
		default:
			if t >= NPCLore0 {
				tileRef += "lore_" + fmt.Sprint(t-NPCLore0)
			} else {
				panic("Unknown npc... " + fmt.Sprint(t))
			}
		}

		DrawTile(screen, tileRef, coords[1], coords[0])
	}

	for str, left := range g.Map.Monsters {
		coords := StringToCoords(str)
		tileRef := "monster"

		if len(left) == 0 {
			tileRef += "_done"
		}

		// TODO: Figure out what to do with dead guys
		// TODO: Draw entities
		DrawTile(screen, tileRef, coords[1], coords[0])
	}
}

func DrawDialogEntities(sprites []string, screen *ebiten.Image) {
	screen.DrawImage(CurrentResourcePack.Background, &ebiten.DrawImageOptions{})
	
	leftOverSpace := SCREEN_DIMENSIONS-(len(sprites)*CurrentResourcePack.SpriteSize*2)

	if leftOverSpace < 0 {
		panic("too many sprites!")
	}

	perSpace := int(math.Round(float64(leftOverSpace)/float64(len(sprites)+1)))

	curX := perSpace

	for _, sprite := range sprites {
		DrawSprite(screen, sprite, curX, 355)
		curX += perSpace
	}
}