package component

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type NonCollisionObjectData struct {
	Position resolv.Vector
}

var NonCollisionObject = donburi.NewComponentType[NonCollisionObjectData]()
