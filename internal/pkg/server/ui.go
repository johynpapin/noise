package server

import (
	"github.com/johynpapin/noise/pkg/things"
)

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var thingsToPositions = make(map[things.Thing]position)

func storePositionOfThing(thing things.Thing, position position) {
	thingsToPositions[thing] = position
}

func getPositionOfThing(thing things.Thing) position {
	return thingsToPositions[thing]
}
