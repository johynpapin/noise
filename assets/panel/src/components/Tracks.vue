<template>
  <div id="tracks">
    <nav class="panel">
      <p class="panel-heading">Tracks</p>
      <div class="panel-block">
        <button class="button is-primary is-fullwidth" @click="openModal">
          <span class="icon">
            <fa icon="plus"></fa>
          </span>
          <span>Create a track</span>
        </button>
      </div>
      <track-item @click.native="selectTrack(track.ID)" v-for="(track, index) in tracks" :key="track.ID" :name="track.name" :selected="track.ID === selectedTrackID"></track-item>
    </nav>
    <div class="modal" :class="[{'is-active': modalActive}]">
      <div class="modal-background" @click="closeModal"></div>
      <div class="modal-content">
        <div class="box">
          <h1 class="title">Create a track</h1>
          <form @submit="createTrack">
            <div class="field">
              <label class="label">Track name</label>
              <div class="control">
                <input class="input" type="text" v-model="trackName" placeholder="Track name">
              </div>
            </div>
            <div class="control">
              <button class="button is-primary">Add track</button>
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

import TrackItem from './TrackItem.vue'

export default {
  name: 'Tracks',
  data () {
    return {
      modalActive: false,
      trackName: ''
    }
  },
  computed: mapState({
    tracks: state => state.state.tracks,
    selectedTrackID: state => state.selectedTrackID
  }),
  methods: {
    openModal () {
      this.trackName = ''
      this.modalActive = true
    },
    closeModal () {
      this.modalActive = false
    },
    createTrack () {
      this.closeModal()
      this.$socket.emit('command', JSON.stringify({command: 'createTrack', name: this.trackName}))
    },
    selectTrack (trackID) {
      this.$store.state.selectedInstrumentPatchID = null
      this.$store.state.selectedTrackID = trackID
      this.$store.state.selectedInstrumentID = null
    },
  },
  components: {
    TrackItem
  }
}
</script>
