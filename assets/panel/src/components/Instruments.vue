<template>
  <div id="instruments">
    <nav class="panel">
      <p class="panel-heading">Instruments</p>
      <div class="panel-block">
        <button class="button is-primary is-fullwidth" @click="openModal">
          <span class="icon">
            <fa icon="plus"></fa>
          </span>
          <span>Create an instrument</span>
        </button>
      </div>
      <instrument-item @click.native="selectInstrument(instrument.ID)" v-for="(instrument, index) in instruments" :key="instrument.ID" :name="instrument.name" :selected="instrument.ID === selectedInstrumentID"></instrument-item>
    </nav>
    <div class="modal" :class="[{'is-active': modalActive}]">
      <div class="modal-background" @click="closeModal"></div>
      <div class="modal-content">
        <div class="box">
          <h1 class="title">Create an instrument</h1>
          <form @submit="createInstrument">
            <div class="field">
              <label class="label">Instrument name</label>
              <div class="control">
                <input class="input" type="text" v-model="instrumentName" placeholder="Instrument name">
              </div>
            </div>
            <div class="control">
              <button class="button is-primary">Create instrument</button>
            </div>
          </form>
        </div>
      </div>
      <button class="modal-close is-large" aria-label="close" @click="closeModal"></button>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

import InstrumentItem from './InstrumentItem.vue'

export default {
  name: 'Tracks',
  data () {
    return {
      modalActive: false,
      instrumentName: ''
    }
  },
  computed: mapState({
    instruments: state => state.state.instruments,
    selectedInstrumentID: state => state.selectedInstrumentID
  }),
  methods: {
    openModal () {
      this.instrumentName = ''
      this.modalActive = true
    },
    closeModal () {
      this.modalActive = false
    },
    createInstrument () {
      this.closeModal()
      this.$socket.emit('command', JSON.stringify({command: 'createInstrument', name: this.instrumentName}))
    },
    selectInstrument (instrumentID) {
      this.$store.state.selectedInstrumentPatchID = null
      this.$store.state.selectedInstrumentID = instrumentID
      this.$store.state.selectedTrackID = null
    },
  },
  components: {
    InstrumentItem
  }
}
</script>
