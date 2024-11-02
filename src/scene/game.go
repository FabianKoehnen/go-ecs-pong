package scene

import (
	"ecs-pong/archetype"
	"ecs-pong/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type GameScene struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

type System interface {
	Update(w donburi.World, screenWidth, screenHeight int)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

func createWorld() donburi.World {
	world := donburi.NewWorld()

	return world
}

func NewGameScene(screenWidth, screenHeight int) *GameScene {
	world := createWorld()
	space := resolv.NewSpace(screenWidth, screenHeight, 1, 1)
	archetype.NewPlayers(world, space, screenWidth, screenHeight)
	archetype.NewBall(world, space, screenWidth, screenHeight)
	return &GameScene{
		world: world,
		systems: []System{
			system.NewPlayerMovement(),
			system.NewVelocity(),
			//system.NewBounce(),
		},
		drawables: []Drawable{
			system.NewRender(),
		},
	}
}

func (t *GameScene) Update(screenWidth, screenHeight int) {
	for _, systemItem := range t.systems {
		systemItem.Update(t.world, screenWidth, screenHeight)
	}
}

func (t *GameScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range t.drawables {
		s.Draw(t.world, screen)
	}
}
