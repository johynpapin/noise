<template>
  <div class="instrument">
    <div class="level">
      <div class="level-item has-text-centered">
        <p class="title">{{name}}</p>
      </div>
    </div>
    <loop v-for="loop in loops" :key="loop.ID" :id="loop.ID" :isRecording="loop.isRecording"></loop>
    <div class="level">
      <div class="level-item has-text-centered">
        <button @click="createLoop" class="button is-small">Create a loop</button>
      </div>
    </div>
  </div>
</template>

<script>
import Loop from './Loop.vue'

export default {
  name: 'InstrumentPatch',
  props: {
    id: Number,
    name: String
  },
  computed: {
    loops () {
      let loops = []

      if (!this.$store.state.state.loops) {
        return loops
      }

      for (let loop of this.$store.state.state.loops) {
        if (loop.instrumentPatchID === this.id) {
          loops.push(loop)
        }
      }

      return loops
    }
  },
  methods: {
    createLoop () {
      this.$socket.emit('command', JSON.stringify({command: 'createLoop', instrumentPatchID: this.id}))
    }
  },
  data () {
    return {
    }
  },
  components: {
    Loop
  }
}
</script>

<style scoped>
.level {
  margin: 0;
  border-bottom: 1px solid;
}
</style>
