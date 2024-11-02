package archetype

import (
	"ecs-pong/assets"
	"ecs-pong/component"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	paddleSizeX = 10
	paddleSizeY = 100
)

func NewPlayers(w donburi.World, space *resolv.Space, initScreenSizeX, initScreenSizeY int) []donburi.Entity {
	removeExistingPlayers(w)
	players := w.CreateMany(2,
		component.Player,
		component.CollisionObject,
		component.Sprite)

	for i, playerEntity := range players {
		entry := w.Entry(playerEntity)
		component.Player.Set(entry, &component.PlayerData{
			LeftPlayer: i == 0,
		})

		img := assets.GetPaddleImage(paddleSizeX, paddleSizeY)
		component.Sprite.Set(entry, &component.SpriteData{
			Image: img,
		})

		if i == 0 {
			component.CollisionObject.Set(entry, component.CreateCollisionObjectData(
				resolv.NewObject(
					0,
					float64(initScreenSizeY/2-paddleSizeY/2),
					float64(img.Bounds().Dx()),
					float64(img.Bounds().Dy()),
				),
				space,
			))
		} else {
			component.CollisionObject.Set(entry, component.CreateCollisionObjectData(
				resolv.NewObject(
					float64(initScreenSizeX-paddleSizeX),
					float64(initScreenSizeY/2-paddleSizeY/2),
					float64(img.Bounds().Dx()),
					float64(img.Bounds().Dy()),
				),
				space,
			))
		}

	}

	return players
}

func removeExistingPlayers(w donburi.World) {
	for entry := range donburi.NewQuery(filter.Contains(component.Player)).Iter(w) {
		entry.Remove()
	}
}
