package system

import (
	"ecs-pong/component"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Render struct {
	query *donburi.Query
}

func NewRender() *Render {
	return &Render{
		query: donburi.NewQuery(
			filter.Contains(
				component.Position,
				component.Sprite,
			))}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	for entry := range r.query.Iter(w) {
		position := component.Position.Get(entry)
		sprite := component.Sprite.Get(entry)

		op := &ebiten.DrawImageOptions{}
		//sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		//op.GeoM.Translate(position.X*sw, position.Y*sh)
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(sprite.Image, op)
	}
}
