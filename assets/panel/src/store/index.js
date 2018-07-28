import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import * as mutations from './mutations'

Vue.use(Vuex)

const state = {
  socket: {
    isConnected: false
  },
  selectedTrackID: null,
  selectedInstrumentID: null,
  selectedInstrumentPatchID: null,
  state: {
    tracks: [],
    instruments: [],
    loops: []
  },
  things: [],
  MIDIEditMode: false
}

const getters = {
  selectedTrack: state => {
    if (state.selectedTrackID === null || !state.state.tracks) {
      return null
    }

    for (let track of state.state.tracks) {
      if (track.ID === state.selectedTrackID) {
        return track
      }
    }

    return null
  },
  selectedInstrument: state => {
    if (state.selectedInstrumentID === null || !state.state.instruments) {
      return null
    }

    for (let instrument of state.state.instruments) {
      if (instrument.ID === state.selectedInstrumentID) {
        return instrument
      }
    }
  },
  selectedInstrumentPatch: (state, getters) => {
    if (state.selectedTrackID === null || !state.state.tracks || state.selectedInstrumentPatchID === null) {
      return null
    }

    for (let patch of getters.selectedTrack.instrumentPatchs) {
      if (patch.ID === state.selectedInstrumentPatchID) {
        return patch
      }
    }
  }
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
