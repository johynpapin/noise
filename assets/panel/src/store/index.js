import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import * as mutations from './mutations'

Vue.use(Vuex)

const state = {
  socket: {
    isConnected: false
  },
  selectedTrackName: null,
  state: {
    tracks: []
  },
  things: [],
  midi: false
}

const getters = {
  selectedTrack: state => {
    if (!state.selectedTrackName || !state.state.tracks) {
      return null
    }

    for (let track of state.state.tracks) {
      if (track.name === state.selectedTrackName) {
        return track
      }
    }

    return null
  }
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
