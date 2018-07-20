package things

type ThingsManager struct {
	namesToThings map[string]Thing
	thingsToNames map[Thing]string
	thingsToKinds map[Thing]string
}

func NewThingsManager() *ThingsManager {
	return &ThingsManager{
		namesToThings: make(map[string]Thing),
		thingsToNames: make(map[Thing]string),
		thingsToKinds: make(map[Thing]string),
	}
}

func (tm *ThingsManager) GetThing(name string) Thing {
	return tm.namesToThings[name]
}

func (tm *ThingsManager) GetName(thing Thing) string {
	return tm.thingsToNames[thing]
}

func (tm *ThingsManager) GetKind(thing Thing) string {
	return tm.thingsToKinds[thing]
}

func (tm *ThingsManager) NewThing(kind string, name string) Thing {
	thing := newThingFromKind(kind)

	tm.namesToThings[name] = thing
	tm.thingsToNames[thing] = name
	tm.thingsToKinds[thing] = kind

	return thing
}
