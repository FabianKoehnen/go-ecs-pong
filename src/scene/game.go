package scene

import (
	"ecs-pong/archetype"
	"ecs-pong/system"
	"github.com/hajimehoshi/ebiten/v2"
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
	archetype.NewPlayers(world, screenWidth, screenHeight)
	archetype.NewBall(world, screenWidth, screenHeight)
	return &GameScene{
		world: world,
		systems: []System{
			system.NewPlayerMovement(),
			system.NewVelocity(),
			system.NewBounce(),
		},
		drawables: []Drawable{
			system.NewRender(),
		},
	}
}

func (t *GameScene) Update(screenWidth, screenHeight int) {
	for _, system := range t.systems {
		system.Update(t.world, screenWidth, screenHeight)
	}
}

func (t *GameScene) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range t.drawables {
		s.Draw(t.world, screen)
	}
}

//func (g *GameScene) Update(screenWidth, screenHeight int) {
//
//}
//func (g *GameScene) Draw(screen *ebiten.Image) {
//
//}

//func NewGame(players []system.ChosenPlayer, screenWidth int, screenHeight int) *GameScene {
//	g := &GameScene{
//		players:      players,
//		level:        0,
//		screenWidth:  screenWidth,
//		screenHeight: screenHeight,
//	}
//
//	g.loadLevel()
//
//	return g
//}
