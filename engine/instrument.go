package engine

import (
	"github.com/johynpapin/noise/things"
)

type Instrument struct {
	ID         int
	Name       string
	EntryPoint *things.Output
	ExitPoint  *things.Output
	Things     []*things.Thing
	Settings   []*things.Setting
}

func (i *Instrument) Next(sample int, entryPoint int) float64 {
	if i.ExitPoint == nil {
		return 0
	}

	return i.ExitPoint.Next(sample, entryPoint)
}

func (i *Instrument) SetEntryPointNextable(entryPointNextable things.Nextable) {
	i.EntryPoint.Nextable = entryPointNextable
}

func (i *Instrument) SetExitPoint(exitPoint *things.Output) {
	i.ExitPoint = exitPoint
}

func (i *Instrument) AddThing(thing *things.Thing) {
	i.Things = append(i.Things, thing)
}

func (i *Instrument) ExposeSetting(name string, setting *things.Setting) {
	setting.ExposedName = name
	i.Settings = append(i.Settings, setting)
}

func (i *Instrument) HideSetting(setting *things.Setting) {
	setting.ExposedName = ""

	var index int
	for i, s := range i.Settings {
		if s == setting {
			index = i
			break
		}
	}

	i.Settings = append(i.Settings[:index], i.Settings[index+1:]...)
}

func NewInstrument(name string) *Instrument {
	return &Instrument{
		ID:         -1,
		Name:       name,
		EntryPoint: things.NewEntryPoint(nil),
	}
}
