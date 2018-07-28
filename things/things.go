package things

type newThingFunc func(string) *Thing

var kindsToThings = make(map[string]newThingFunc)
var Kinds []string

func registerThing(kind string, newThing newThingFunc) {
	kindsToThings[kind] = newThing
	Kinds = append(Kinds, kind)
}

func NewThingFromKind(name string, kind string) *Thing {
	return kindsToThings[kind](name)

}
