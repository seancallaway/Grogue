package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Levels       []Level
	CurrentLevel *Level
	Entities     []Entity
	Player       *Entity
}

// Creates a new Game object and initializes the data.
func NewGame() *Game {
	g := &Game{}
	g.Levels = append(g.Levels, NewLevel())
	g.CurrentLevel = &g.Levels[0]

	startX, startY := g.CurrentLevel.Rooms[0].Center()
	player, err := NewEntity(startX, startY, "player")
	if err != nil {
		log.Fatal(err)
	}
	g.Entities = append(g.Entities, player)
	g.Player = &g.Entities[0]
	return g
}

func (g *Game) Update() error {
	HandleInput(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Levels[0]
	level.Draw(screen)
	RenderEntities(g, level, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowTitle("Grogue")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
