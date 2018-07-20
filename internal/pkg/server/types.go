package server

type stateMessage struct {
	State state `json:"state"`
}

type thingsMessage struct {
	Things []string `json:"things"`
}

// General commands

type createTrackMessage struct {
	Name string `json:"name"`
}

type deleteTrackMessage struct {
	Name string `json:"name"`
}

type createThingMessage struct {
	TrackName string `json:"trackName"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type deleteThingMessage struct {
	TrackName string `json:"trackName"`
	Name      string `json:"name"`
}

type attachToThingMessage struct {
	TrackName       string `json:"trackName"`
	InputThingName  string `json:"inputThingName"`
	InputName       string `json:"inputName"`
	OutputThingName string `json:"outputThingName"`
	OutputName      string `json:"outputName"`
}

type detachFromThingMessage struct {
	TrackName      string `json:"trackName"`
	InputThingName string `json:"inputThingName"`
	InputName      string `json:"inputName"`
}

type attachToTrackMessage struct {
	TrackName string `json:"trackName"`
	ThingName string `json:"thingName"`
	Name      string `json:"name"`
}

type detachFromTrackMessage struct {
	TrackName string `json:"trackName"`
	ThingName string `json:"thingName"`
	Name      string `json:"Name"`
}

type updateSettingMessage struct {
	TrackName string      `json:"trackName"`
	ThingName string      `json:"thingName"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
}

// MIDI commands

type attachToMIDIMessage struct {
	TrackName string `json:"trackName"`
	ThingName string `json:"thingName"`
	Name      string `json:"name"`
}

type detachFromMIDIMessage struct {
	TrackName string `json:"trackName"`
	ThingName string `json:"thingName"`
	Name      string `json:"name"`
}
