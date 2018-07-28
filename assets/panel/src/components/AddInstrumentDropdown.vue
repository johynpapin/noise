<template>
  <div class="add-instrument-dropdown">
    <div class="dropdown is-hoverable">
      <div class="dropdown-trigger">
        <button class="button is-fullwidth is-primary is-inverted" aria-haspopup="true" aria-controls="add-instrument-menu">
          <span class="icon">
            <fa icon="plus"></fa>
          </span>
          <span>Add an instrument</span>
        </button>
        <div class="dropdown-menu" id="add-instrument-menu" role="menu">
          <div class="dropdown-content">
            <a @click="addInstrument(instrument)" v-for="instrument in instruments" :key="instrument.ID" class="dropdown-item">
              {{instrument.name}}
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'AddInstrumentDropdown',
  props: {
    track: Object
  },
  data () {
    return {
    }
  },
  computed: mapState({
    instruments: state => state.state.instruments
  }),
  methods: {
    addInstrument (instrument) {
      console.log(instrument, this.track)
      this.$socket.emit('command', JSON.stringify({command: 'addInstrumentToTrack', trackID: this.track.ID, instrumentID: instrument.ID}))
    }
  }
}
</script>
