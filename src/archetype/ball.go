package archetype

import (
	"ecs-pong/assets"
	"ecs-pong/component"
	"ecs-pong/util"
	"github.com/yohamta/donburi"
)

const (
	size = 20

	velocityMax = 2
	velocityMin = -2
)

func NewBall(w donburi.World, screenSizeX, screenSizeY int) {
	ball := w.Entry(w.Create(
		component.Sprite,
		component.Position,
		component.Velocity,
		component.Bouncing,
	))

	img := assets.GetBallImage(size)
	component.Sprite.Set(ball, &component.SpriteData{Image: img})
	component.Position.Set(ball, &component.PositionData{
		X: float64(screenSizeX/2 - img.Bounds().Dx()/2),
		Y: float64((screenSizeY / 2) - img.Bounds().Dy()/2),
	})
	component.Velocity.Set(ball, &component.VelocityData{
		X: util.RandomFloat(velocityMin, velocityMax),
		Y: util.RandomFloat(velocityMin, velocityMax),
	})
	component.Bouncing.Set(ball, &component.BouncingData{})
}
