<template>
  <div class="track-screen">
    <div class="instruments">
      <div class="instrument-column" v-for="patch in track.instrumentPatchs" :key="patch.ID">
        <instrument-patch @click.native="selectInstrumentPatch(patch)" :id="patch.ID" :name="patch.instrument.name"></instrument-patch>
      </div>
      <div class="instrument-column">
        <div class="level">
          <div class="level-item has-text-centered">
            <add-instrument-dropdown :track="track"></add-instrument-dropdown>
          </div>
        </div>
      </div>
    </div>
    <div class="instrument-patch-settings-pane">
      <instrument-patch-settings v-if="selectedInstrumentPatch !== null" :patch="selectedInstrumentPatch"></instrument-patch-settings>
    </div>
  </div>
</template>

<script>
import InstrumentPatch from './InstrumentPatch.vue'
import InstrumentPatchSettings from './InstrumentPatchSettings.vue'
import AddInstrumentDropdown from './AddInstrumentDropdown.vue'

import { mapGetters } from 'vuex'

export default {
  name: 'TrackScreen',
  props: {
    track: Object
  },
  data () {
    return {
    }
  },
  computed: mapGetters([
    'selectedInstrumentPatch'
  ]),
  methods: {
    selectInstrumentPatch(patch) {
      this.$socket.emit('command', JSON.stringify({command: 'updateCurrentInstrumentPatch', instrumentPatchID: patch.ID}))
      this.$store.state.selectedInstrumentPatchID = patch.ID
    }
  },
  components: {
    InstrumentPatch,
    InstrumentPatchSettings,
    AddInstrumentDropdown
  }
}
</script>

<style scoped>
.track-screen {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.instruments {
  flex: 1 1 auto;
  display: flex;
}

.instrument-patch-settings-pane {
  height: 300px;
  border-top: 1px solid;
}

.instrument-column {
  height: 100%;

  flex: 0 0 200px;

  border-right: 1px solid;

  display: inline-block;
}
</style>
