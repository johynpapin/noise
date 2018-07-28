<template>
  <div id="home">
    <div class="columns is-gapless">
      <div class="column is-one-quarter">
        <sidebar></sidebar>
      </div>
      <div class="column">
        <track-screen v-if="selectedTrack !== null" :track="selectedTrack"></track-screen>
        <instrument-editor v-else-if="selectedInstrument !== null" :instrument="selectedInstrument"></instrument-editor>
        <div v-else class="container">
          <h1 class="title is-1 is-spaced">No track selected.</h1>
          <h3 class="subtitle is-3">NOISE cannot play music.</h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TrackScreen from './TrackScreen.vue'
import InstrumentEditor from './InstrumentEditor.vue'
import Sidebar from './Sidebar.vue'

import { mapGetters } from 'vuex'

export default {
  name: 'Home',
  data () {
    return {
    }
  },
  computed: mapGetters([
    'selectedTrack',
    'selectedInstrument'
  ]),
  components: {
    TrackScreen,
    InstrumentEditor,
    Sidebar
  }
}
</script>

<style scoped>
#home, #home .columns, #home .column {
  height: 100%;
}
</style>
