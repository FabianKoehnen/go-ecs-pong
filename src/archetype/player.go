package archetype

import (
	"ecs-pong/assets"
	"ecs-pong/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	paddleSizeX = 10
	paddleSizeY = 100
)

func NewPlayers(w donburi.World, initScreenSizeX, initScreenSizeY int) []donburi.Entity {
	removeExistingPlayers(w)
	players := w.CreateMany(2,
		component.Player,
		component.Position,
		component.Sprite)

	for i, playerEntity := range players {
		entry := w.Entry(playerEntity)
		component.Player.Set(entry, &component.PlayerData{
			LeftPlayer: i == 0,
		})
		if i == 0 {

			component.Position.Set(entry, &component.PositionData{
				X: 0,
				Y: float64(initScreenSizeY/2 - paddleSizeY/2),
			})
		} else {
			component.Position.Set(entry, &component.PositionData{
				X: float64(initScreenSizeX - paddleSizeX),
				Y: float64(initScreenSizeY/2 - paddleSizeY/2),
			})
		}
		component.Sprite.Set(entry, &component.SpriteData{
			Image: assets.GetPaddleImage(paddleSizeX, paddleSizeY),
		})
	}

	return players
}

func removeExistingPlayers(w donburi.World) {
	for entry := range donburi.NewQuery(filter.Contains(component.Player)).Iter(w) {
		entry.Remove()
	}
}
