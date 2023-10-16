package main

import "github.com/hajimehoshi/ebiten/v2"

// Handle user input, including moving the player.
func HandleInput(g *Game) {
	dx := 0
	dy := 0

	// Player Movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dy = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dx = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		dx = 1
	}

	newPos := GetIndexFromCoords(g.Player.X+dx, g.Player.Y+dy)
	tile := g.CurrentLevel.Tiles[newPos]
	if !tile.Blocked {
		g.Player.X += dx
		g.Player.Y += dy
		g.CurrentLevel.PlayerView.Compute(g.CurrentLevel, g.Player.X, g.Player.Y, 8)
	}
}
