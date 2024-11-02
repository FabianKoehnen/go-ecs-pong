package system

import (
	"ecs-pong/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const speed = 8

type PlayerMovement struct {
	query *donburi.Query
}

func NewPlayerMovement() *PlayerMovement {
	return &PlayerMovement{
		query: donburi.NewQuery(
			filter.Contains(
				component.Player,
				component.CollisionObject,
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
			continue
		}

		player := component.CollisionObject.Get(entry).Object
		playerImage := component.Sprite.Get(entry).Image

		if upPressed {
			if player.Position.Y-speed > 0 {
				player.Position.Y -= speed
			} else {
				player.Position.Y = 0
			}
		}
		if downPressed {
			if player.Position.Y+float64(playerImage.Bounds().Dy()) < float64(screenHeight) {
				player.Position.Y += speed
			} else {
				player.Position.Y = float64(screenHeight - playerImage.Bounds().Dy())
			}
		}

		player.Update()
	}
}
