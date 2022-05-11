package main

func (g *Game) ArrowKey(dirs map[Direction]bool) *APIError {
	if dirs[DirDown] && dirs[DirUp] {
		delete(dirs, DirDown)
		delete(dirs, DirUp)
	}

	if dirs[DirLeft] && dirs[DirRight] {
		delete(dirs, DirLeft)
		delete(dirs, DirRight)
	}

	var err *APIError

	for dir, p := range dirs {
		if !p {
			continue
		}

		newY := g.Map.User[0] + DirectionEnum[dir][0]
		newX := g.Map.User[1] + DirectionEnum[dir][1]

		if !LocOk(newY) || !LocOk(newX) {
			continue
		}

		g.Direction = dir

		switch g.Map.Layout[newY][newX] {
		case MapArena, MapEmptySpace:
			err = g.Map.Move(dir, 1)
		default:
			continue
		}
	
		if err != nil {
			return err
		}

	}

	return nil
}

func (g *Game) ZKey() *APIError {
	var err *APIError

	switch g.Screen {
	case S_Tile:
		c := Coords{g.Map.User[0] + DirectionEnum[g.Direction][0], g.Map.User[1] + DirectionEnum[g.Direction][1]}
		tile := g.Map.Layout[c[0]][c[1]]

		switch tile {
			case MapExit:
				err = g.Map.Action(g.Direction)
			case MapChest:
				// TODO:
			case MapNPC:
				if g.Map.NPC[c.String()] >= NPCLore0 {
					g.SetScreen(S_Lore)
					g.NPC = NPCProg{
						Text: []string{"Hello there! I am very angry. I hate you in fact.", 
										"AHAHAHAHAHHAHAHAHAHAHAHA PRANKED",
						},
					}
				}
		}
	}

	if err != nil {
		return err
	}

	return nil
}