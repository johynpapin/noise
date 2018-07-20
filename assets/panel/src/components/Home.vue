<template>
  <div id="home" class="columns">
    <div class="column is-one-quarter">
      <tracks></tracks>
    </div>
    <div class="column">
      <track-screen v-if="selectedTrack !== null" :track="selectedTrack"></track-screen>
      <div v-else class="container">
        <h1 class="title is-1 is-spaced">No track selected.</h1>
        <h3 class="subtitle is-3">NOISE cannot play music.</h3>
      </div>
    </div>
  </div>
</template>

<script>
import Tracks from './Tracks.vue'
import TrackScreen from './TrackScreen.vue'

import { mapGetters } from 'vuex'

export default {
  name: 'Home',
  data () {
    return {
    }
  },
  computed: mapGetters([
    'selectedTrack'
  ]),
  components: {
    Tracks,
    TrackScreen
  }
}
</script>

<style scoped>
#home, #home .column {
  height: 100%;
}
</style>
