package archetype

import (
	"ecs-pong/assets"
	"ecs-pong/component"
	"ecs-pong/util"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

const (
	size = 20

	velocityMax = 5
)

func NewBall(w donburi.World, space *resolv.Space, screenSizeX, screenSizeY int) {
	ball := w.Entry(w.Create(
		component.Sprite,
		component.CollisionObject,
		component.Velocity,
		component.Bouncy,
	))

	img := assets.GetBallImage(size)
	component.Sprite.Set(ball, &component.SpriteData{Image: img})
	component.CollisionObject.Set(ball, component.CreateCollisionObjectData(
		resolv.NewObject(
			float64(screenSizeX/2-img.Bounds().Dx()/2),
			float64((screenSizeY/2)-img.Bounds().Dy()/2),
			float64(img.Bounds().Dx()),
			float64(img.Bounds().Dy()),
		),
		space,
	))
	component.Velocity.Set(ball, &resolv.Vector{
		X: util.RandomFloat(velocityMax, velocityMax),
		Y: util.RandomFloat(velocityMax, velocityMax),
	})
}
