package server

import (
	"github.com/johynpapin/noise/store"
)

type stateMessage struct {
	State *store.State `json:"state"`
}

type thingsMessage struct {
	Things []string `json:"things"`
}

type commandMessage struct {
	Command string `json:"command"`

	TrackID           int `json:"trackID"`
	InstrumentPatchID int `json:"instrumentPatchID"`
	InstrumentID      int `json:"instrumentID"`
	ThingID           int `json:"thingID"`
	SettingID         int `json:"settingID"`
	InputID           int `json:"inputID"`
	OutputID          int `json:"outputID"`
	LoopID            int `json:"loopID"`

	Name  string  `json:"name"`
	Kind  string  `json:"kind"`
	Value float64 `json:"value"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`

	X float64 `json:"x"`
	Y float64 `json:"y"`
}
