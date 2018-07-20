export const SOCKET_CONNECT = (state, status) => {
  state.socket.isConnected = true
}

export const SOCKET_DISCONNECT = (state, status) => {
  state.socket.isConnected = false
}

export const SOCKET_STATE = (state, message) => {
  state.state = JSON.parse(message).state
}

export const SOCKET_THINGS = (state, message) => {
  state.things = JSON.parse(message).things
}
