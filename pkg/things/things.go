package things

var kindsToThings = make(map[string]func() Thing)
var Kinds []string

func registerThing(kind string, newThing func() Thing) {
	kindsToThings[kind] = newThing
	Kinds = append(Kinds, kind)
}

func newThingFromKind(kind string) Thing {
	return kindsToThings[kind]()
}
