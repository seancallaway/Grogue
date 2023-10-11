package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Entity struct {
	X     int
	Y     int
	Image *ebiten.Image
}

// Create an Entity object at tile coordinates 'x, y' represented by a PNG named 'imageName'.
func NewEntity(x int, y int, imageName string) (Entity, error) {
	image, _, err := ebitenutil.NewImageFromFile("assets/" + imageName + ".png")
	if err != nil {
		return Entity{}, err
	}

	entity := Entity{
		X:     x,
		Y:     y,
		Image: image,
	}
	return entity, nil
}

// Move the entity by a given amount.
func (entity *Entity) Move(dx int, dy int) {
	entity.X += dx
	entity.Y += dy
}
