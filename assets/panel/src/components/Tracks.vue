<template>
  <div id="tracks">
    <nav class="panel">
      <p class="panel-heading">Tracks</p>
      <div class="panel-block">
        <button class="button is-primary is-fullwidth" @click="openAddModal">
          <span class="icon">
            <fa icon="plus"></fa>
          </span>
          <span>Add a track</span>
        </button>
      </div>
      <track-item @click.native="selectTrack(index)" v-for="(track, index) in tracks" :key="track.name" :name="track.name" :selected="index === selectedTrackIndex"></track-item>
    </nav>
    <div class="modal" :class="[{'is-active': addModalActive}]">
      <div class="modal-background" @click="closeAddModal"></div>
      <div class="modal-content">
        <div class="box">
          <h1 class="title">Add a track</h1>
          <form @submit="addTrack">
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
      <button class="modal-close is-large" aria-label="close" @click="closeAddModal"></button>
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
      addModalActive: false,
      trackName: ''
    }
  },
  computed: mapState({
    tracks: state => state.state.tracks,
    selectedTrackIndex: state => state.selectedTrackIndex
  }),
  methods: {
    openAddModal () {
      this.trackName = ''
      this.addModalActive = true
    },
    closeAddModal () {
      this.addModalActive = false
    },
    addTrack () {
      this.closeAddModal()
      this.$socket.emit('addTrack', JSON.stringify({name: this.trackName}))
    },
    selectTrack (trackIndex) {
      this.$store.state.selectedTrackIndex = trackIndex
    }
  },
  components: {
    TrackItem
  }
}
</script>
