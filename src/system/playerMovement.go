package system

import (
	"ecs-pong/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const speed = 5

type PlayerMovement struct {
	query *donburi.Query
}

func NewPlayerMovement() *PlayerMovement {
	return &PlayerMovement{
		query: donburi.NewQuery(
			filter.Contains(
				component.Player,
				component.Position,
				component.Sprite,
			)),
	}
}

func (p *PlayerMovement) Update(w donburi.World, screenWidth, screenHeight int) {
	for entry := range p.query.Iter(w) {
		upPressed := false
		downPressed := false
		if component.Player.Get(entry).LeftPlayer {
			if ebiten.IsKeyPressed(ebiten.KeyW) {
				upPressed = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyS) {
				downPressed = true
			}
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				upPressed = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) {
				downPressed = true
			}
		}

		if upPressed && downPressed {
			return
		}

		position := component.Position.Get(entry)
		playerImage := component.Sprite.Get(entry).Image

		if upPressed {
			if position.Y-speed > 0 {
				position.Y -= speed
			} else {
				position.Y = 0
			}
		}
		if downPressed {
			if position.Y+float64(playerImage.Bounds().Dy()) < float64(screenHeight) {
				position.Y += speed
			} else {
				position.Y = float64(screenHeight - playerImage.Bounds().Dy())
			}
		}
	}
}
