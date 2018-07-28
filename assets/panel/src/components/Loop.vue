<template>
  <div class="level">
    <div class="level-item has-text-centered">
      <button @click="stopRecording" v-if="isRecording" class="button is-small">Stop recording</button>
      <button @click="startRecording" v-else class="button is-small">Start recording</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Loop',
  props: {
    id: Number,
    isRecording: Boolean
  },
  data () {
    return {
    }
  },
  methods: {
    startRecording () {
      this.$socket.emit('command', JSON.stringify({command: 'startRecording', loopID: this.id}))
    },
    stopRecording () {
      this.$socket.emit('command', JSON.stringify({command: 'stopRecording', loopID: this.id}))
    }
  }
}
</script>
