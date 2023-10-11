package main

import "github.com/hajimehoshi/ebiten/v2"

// Renders all of the entities in a given game onto screen.
func RenderEntities(g *Game, level Level, screen *ebiten.Image) {
	for _, entity := range g.Entities {
		idx := GetIndexFromCoords(entity.X, entity.Y)
		tile := level.Tiles[idx]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(entity.Image, op)
	}
}
