package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// Creates a new instance of the static game data.
func NewGameData() GameData {
	gd := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
	return gd
}

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Opaque  bool
	Image   *ebiten.Image
}

const (
	TileFloor string = "floor"
	TileWall  string = "wall"
)

// Creates a MapTile of a given type at pixels x, y.
//
//	Supported types include 'Floor' and 'Wall'
func NewTile(x int, y int, tileType string) (MapTile, error) {
	blocked := true
	opaque := true

	image, _, err := ebitenutil.NewImageFromFile("assets/" + tileType + ".png")
	if err != nil {
		return MapTile{}, err
	}

	if tileType == TileFloor {
		blocked = false
		opaque = false
	}

	tile := MapTile{
		PixelX:  x,
		PixelY:  y,
		Blocked: blocked,
		Opaque:  opaque,
		Image:   image,
	}
	return tile, nil
}

// Returns the logical index of the map slice from given X and Y tile coordinates.
func GetIndexFromCoords(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

type Level struct {
	Tiles []MapTile
	Rooms []RectangularRoom
}

// Creates a new Level object
func NewLevel() Level {
	l := Level{}
	l.createTiles()
	return l
}

// Creates a map of tiles.
func (level *Level) createTiles() {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)

	// Fill with wall tiles
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			idx := GetIndexFromCoords(x, y)
			wall, err := NewTile(x*gd.TileWidth, y*gd.TileHeight, TileWall)
			if err != nil {
				log.Fatal(err)
			}
			tiles[idx] = wall
		}
	}

	room1 := NewRectangularRoom(25, 15, 10, 15)
	room2 := NewRectangularRoom(40, 15, 10, 15)
	level.Rooms = append(level.Rooms, room1, room2)

	for _, room := range level.Rooms {
		x1, x2, y1, y2 := room.Interior()
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				idx := GetIndexFromCoords(x, y)
				floor, err := NewTile(x*gd.TileWidth, y*gd.TileHeight, TileFloor)
				if err != nil {
					log.Fatal(err)
				}
				tiles[idx] = floor
			}
		}
	}

	level.Tiles = tiles
}

// Draw the current level to 'screen'.
func (level *Level) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[GetIndexFromCoords(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

type RectangularRoom struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Create a new RectangularRoom structure.
func NewRectangularRoom(x int, y int, width int, height int) RectangularRoom {
	return RectangularRoom{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

// Returns the tile coordinates of the center of the RectangularRoom.
func (r *RectangularRoom) Center() (int, int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2
	return centerX, centerY
}

// Returns the tile coordinates of the interior of the RectangularRoom.
func (r *RectangularRoom) Interior() (int, int, int, int) {
	return r.X1 + 1, r.X2 - 1, r.Y1 + 1, r.Y2 - 1
}
