package component

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type CollisionObjectData struct {
	Object *resolv.Object
	Space  *resolv.Space
}

func CreateCollisionObjectData(Object *resolv.Object, Space *resolv.Space) *CollisionObjectData {
	Space.Add(Object)
	return &CollisionObjectData{Object, Space}
}

var CollisionObject = donburi.NewComponentType[CollisionObjectData]()
