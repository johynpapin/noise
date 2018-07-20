import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import * as mutations from './mutations'

Vue.use(Vuex)

const state = {
  socket: {
    isConnected: false
  },
  selectedTrackIndex: null,
  state: {
    tracks: []
  },
  things: []
}

const getters = {
  selectedTrack: state => {
    return state.selectedTrackIndex !== null ? state.state.tracks[state.selectedTrackIndex] : null
  }
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
