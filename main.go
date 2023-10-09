package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Levels []Level
}

// Creates a new Game object and initializes the data.
func NewGame() *Game {
	g := &Game{}
	g.Levels = append(g.Levels, NewLevel())
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Levels[0]
	level.Draw(screen)
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
