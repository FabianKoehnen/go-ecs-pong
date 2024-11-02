package system

import (
	"ecs-pong/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Render struct {
	query *donburi.Query
}

func NewRender() *Render {
	return &Render{
		query: donburi.NewQuery(
			filter.And(
				filter.Contains(component.Sprite),
				filter.Or(
					filter.Contains(component.NonCollisionObject),
					filter.Contains(component.CollisionObject),
				),
			)),
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	for entry := range r.query.Iter(w) {
		var position resolv.Vector
		if entry.HasComponent(component.NonCollisionObject) {
			position = component.NonCollisionObject.Get(entry).Position
		} else if entry.HasComponent(component.CollisionObject) {
			position = component.CollisionObject.Get(entry).Object.Position
		}
		sprite := component.Sprite.Get(entry)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(sprite.Image, op)
	}
}
