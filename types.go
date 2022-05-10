package main

const DIMENSIONS = 26

type MonsterType int8

const (
	MTRange MonsterType = iota
	MTTank
	MTGlassSword
	MTFire
	MTEarth
	MTAir
	MTWater
	MTDark
	MTLight
	MTSlow
	MonsterType10 //TODO:
	MonsterType11
	MonsterType12
)

type Monsters []MonsterType

type Layout [][]TileType

type ChestInfo struct {
	Rare bool `json:"rare"`
	Open bool `json:"open"`
}

type Map struct {
	Layout Layout `json:"layout"`
	Exit   Coords `json:"exit"`
	// strings.Join(Coords, ",") : Monsters
	Monsters map[string]Monsters `json:"monsters"`
	// strings.Join(Coords, ",") : completed them
	Rooms map[string]bool `json:"roomsCleared"`

	// strings.Join(Coords, ",") : completed them
	ChestsInfo map[string]ChestInfo `json:"chests"`

	User  Coords             `json:"user"`
	Level int                `json:"level"`
	NPC   map[string]NPCType `json:"npcs"`
}

type NPCType int8

// 0 - y, 1 - x
type Coords [2]int

type TileType int8

const (
	MapWall TileType = iota // MUST BE THERE BY DEFAULT!
	MapEmptySpace
	MapChest
	MapMonster
	MapArena
	MapNPC
	MapExit
)

const (
	NPCShopWeapon NPCType = iota
	NPCShopArmor
	NPCShopPotion
	NPCUpgrade
	NPCLore0
	NPCLore1
	NPCLore2
	NPCLore3
	NPCLore4
	NPCLore5
	NPCLore6
	NPCLore7
	NPCLore8
	NPCLore9
)

func (tile TileType) IsWalkable() bool {
	return tile == MapArena || tile == MapEmptySpace
}
