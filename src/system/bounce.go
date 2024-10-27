package system

import (
	"ecs-pong/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Bounce struct {
	query       *donburi.Query
	paddelQuery *donburi.Query
}

func NewBounce() *Bounce {
	return &Bounce{
		query: donburi.NewQuery(
			filter.Contains(
				component.Position,
				component.Velocity,
				component.Bouncing,
				component.Sprite,
			)),
		paddelQuery: donburi.NewQuery(
			filter.Contains(
				component.Position,
				component.Sprite,
				component.Player,
			)),
	}
}

func (b *Bounce) Update(w donburi.World, screenWidth, screenHeight int) {
	for entry := range b.query.Iter(w) {
		position := component.Position.Get(entry)
		velocity := component.Velocity.Get(entry)
		sprite := component.Sprite.Get(entry)

		if position.Y <= 0 ||
			position.Y >= float64(screenHeight-sprite.Image.Bounds().Dy()) {
			velocity.Y *= -1
		}

		//for paddelEntry := range b.paddelQuery.Iter(w) {
		//
		//}
	}
}
