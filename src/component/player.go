package component

import (
	"github.com/yohamta/donburi"
)

type PlayerData struct {
	LeftPlayer bool
}

var Player = donburi.NewComponentType[PlayerData]()
