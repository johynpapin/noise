package server

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/runtime"
)

func emitState(so socketio.Socket, r *runtime.Runtime) error {
	state := r.Store.ExportState()
	stateMsg := stateMessage{state}
	answer, _ := json.Marshal(stateMsg)

	return so.Emit("state", string(answer))
}
