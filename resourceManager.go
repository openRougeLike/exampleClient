package main

import (
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type ResourceManifest struct {
	Name    string       `json:"name"`
	Size    int          `json:"size"`
	Tiles   ResourceList `json:"tiles"`
	Sprites ResourceList `json:"sprites"`
	Backgrounds ResourceList `json:"backgrounds"`
	Font 	string `json:"font"`
}

type ResourceList struct {
	List []ResourceTile
	Size int `json:"size"`
}

type ResourceTile struct {
	Max  int    `json:"max"`
	Name string `json:"name"`
}

type ResourcePack struct {
	Tiles      *ebiten.Image
	TileSize   int
	Sprites    *ebiten.Image
	SpriteSize int
	Background *ebiten.Image
	Font 	   font.Face
}

var CurrentResourcePack ResourcePack

func LoadResourcePack() {
	CurrentResourcePack = FetchResourcesPack("minimalist", 0)
}

func FetchResourcesPack(setName string, level int) ResourcePack {
	fManifest, err := ioutil.ReadFile(filepath.Join("tiles", setName, "manifest.json"))
	PanicIfErr(err)

	manifest := ResourceManifest{}

	err = json.Unmarshal(fManifest, &manifest)
	PanicIfErr(err)

	pack := ResourcePack{}

	sprites, sizeSprites := manifest.Sprites.FetchImg(setName, level)
	tiles, sizeTiles := manifest.Tiles.FetchImg(setName, level)

	pack.Sprites = sprites
	pack.Tiles = tiles

	pack.SpriteSize = sizeSprites
	pack.TileSize = sizeTiles

	bg, _ := manifest.Backgrounds.FetchImg(setName, level)

	pack.Background = bg

	fontRaw, err := ioutil.ReadFile(filepath.Join("tiles", setName, manifest.Font))
	PanicIfErr(err)

	fontParsed, err := opentype.Parse(fontRaw)
	PanicIfErr(err)

	face, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	PanicIfErr(err)

	pack.Font = face
	
	return pack
}

func (l ResourceList) FetchImg(setName string, level int) (*ebiten.Image, int) {
	sort.Slice(l.List, func(i, j int) bool {
		if l.List[j].Max == -1 || l.List[i].Max == -1 {
			return true
		}
		return l.List[i].Max < l.List[j].Max
	})

	for _, t := range l.List {
		if t.Max == -1 || level < t.Max {
			f, err := os.Open(filepath.Join("tiles", setName, t.Name))
			PanicIfErr(err)
			img, _, err := image.Decode(f)
			PanicIfErr(err)
			imgE := ebiten.NewImageFromImage(img)
			return imgE, l.Size
		}
	}

	panic("Tileset not found!")
}
