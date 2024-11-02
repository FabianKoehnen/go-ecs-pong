package system

import (
	"ecs-pong/component"
	"fmt"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Velocity struct {
	query *donburi.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		query: donburi.NewQuery(
			filter.And(
				filter.Contains(component.Velocity),
				filter.Or(
					filter.Contains(component.NonCollisionObject),
					filter.Contains(component.CollisionObject),
				)),
		),
	}
}

func (v *Velocity) Update(w donburi.World, screenWidth, screenHeight int) {
	for entry := range v.query.Iter(w) {

		var position *resolv.Vector
		if entry.HasComponent(component.NonCollisionObject) {
			position = &component.NonCollisionObject.Get(entry).Position
		} else if entry.HasComponent(component.CollisionObject) {
			position = &component.CollisionObject.Get(entry).Object.Position
		}
		if position == nil {
			continue
		}

		velocity := component.Velocity.Get(entry)

		if entry.HasComponent(component.CollisionObject) {
			collisionObjEntry := component.CollisionObject.Get(entry)
			collisionObj := collisionObjEntry.Object

			pos := collisionObj.Position

			if pos.X < 0 || pos.X+collisionObj.Size.X > float64(screenWidth) {
				fmt.Println("dead")
				velocity.Set(0, 0)
			}

			if pos.Y < 0 || pos.Y+collisionObj.Size.Y > float64(screenHeight) {
				velocity.Y = -velocity.Y
			}

			if collision := collisionObj.Check(velocity.X, velocity.Y); collision != nil {
				var closestCollisionDelta resolv.Vector

				for _, collidingObj := range collision.Objects {
					collisionDelta := collision.ContactWithObject(collidingObj)
					closestCollisionDelta.X = min(closestCollisionDelta.X, collisionDelta.X)
					closestCollisionDelta.Y = min(closestCollisionDelta.Y, collisionDelta.Y)
				}

				if entry.HasComponent(component.Bouncy) {
					velocity.X = -velocity.X
				} else {
					velocity.X = closestCollisionDelta.X
					velocity.Y = closestCollisionDelta.Y
				}
			}
		}

		position.X += velocity.X
		position.Y += velocity.Y
	}
}
