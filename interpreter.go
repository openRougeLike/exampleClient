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

		switch g.Map.Layout[newY][newX] {
		case MapArena, MapEmptySpace:
			err = g.Map.Move(dir, 1)
		case MapExit:
			err = g.Map.Action(dir)
		case MapChest:
			// TODO:
		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}
