package component

import "github.com/yohamta/donburi"

type BouncingData struct {
}

var Bouncing = donburi.NewComponentType[BouncingData]()
