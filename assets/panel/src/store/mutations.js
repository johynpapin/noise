export const SOCKET_CONNECT = (state, status) => {
  state.socket.isConnected = true
}

export const SOCKET_DISCONNECT = (state, status) => {
  state.socket.isConnected = false
}

export const SOCKET_STATE = (state, message) => {
  let newState = JSON.parse(message).state

  state.state = newState
}

export const SOCKET_THINGS = (state, message) => {
  state.things = JSON.parse(message).things
}
