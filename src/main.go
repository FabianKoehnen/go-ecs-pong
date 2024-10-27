package main

import (
	"ecs-pong/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 1080
	screenHeight = 720
)

type Scene interface {
	Update(screenWidth, screenHeight int)
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene                     Scene
	screenWidth, screenHeight int
}

func NewGame() *Game {
	g := &Game{}
	g.screenWidth, g.screenHeight = screenWidth, screenHeight
	//g.scene = scene.NewGame(screenWidth, screenHeight)
	g.scene = scene.NewGameScene(screenWidth, screenHeight)
	//g.switchToTitle()
	return g
}

func (g *Game) Update() error {
	g.scene.Update(g.screenWidth, g.screenHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//rand.Seed(time.Now().UTC().UnixNano())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
