package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func GetPaddleImage(sizeX, sizeY int) *ebiten.Image {
	img := ebiten.NewImage(sizeX, sizeY)

	img.Fill(color.White)

	return img
}

func GetBallImage(size int) *ebiten.Image {
	img := ebiten.NewImage(size, size)

	img.Fill(color.White)

	return img
}
