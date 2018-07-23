<template>
  <div id="track-screen">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <div class="columns">
            <div class="column">
                <h1 class="title">{{track.name}}</h1>
                <h2 class="subtitle">Such a perfect track</h2>
            </div>
            <div class="column is-narrow">
              <add-thing-dropdown :track="track"></add-thing-dropdown>
            </div>
          </div>
        </div>
      </div>
    </section>
    <blocks-container @drop-block="dropThing" @click-setting="clickSetting" @update-setting="updateSetting" @new-link="newLink" :raw-blocks="blocks"></blocks-container>
  </div>
</template>

<script>
import BlocksContainer from './blocks/BlocksContainer.vue'
import AddThingDropdown from './AddThingDropdown.vue'

export default {
  name: 'TrackScreen',
  props: {
    track: Object
  },
  computed: {
    blocks () {
      let blocks = [{
        name: 'Output',
        inputs: [{name: 'main'}],
        x: 0,
        y: 0,
        outputs: []
      }]

      if (!this.track.things) {
        return blocks
      }

      if (this.track.outputs) {
        for (let output of this.track.outputs) {
          blocks[0].inputs[0].linkedTo = {blockName: output.thingName, outputName: output.name}
        }
      }

      for (let block of this.track.things) {
        for (let inputIndex in block.inputs) {
          if (block.inputs[inputIndex].linkedTo) {
            block.inputs[inputIndex].linkedTo.blockName = block.inputs[inputIndex].linkedTo.thingName
          }
        }

        blocks.push(block)
      }

      return blocks
    }
  },
  data () {
    return {
    }
  },
  methods: {
    newLink (link) {
      if (link.inputBlockName === 'Output') {
        this.$socket.emit('attachToTrack', JSON.stringify({trackName: this.track.name, thingName: link.outputBlockName, name: link.outputName}))
      } else {
        this.$socket.emit('attachToThing', JSON.stringify({trackName: this.track.name, outputThingName: link.outputBlockName, inputThingName: link.inputBlockName, outputName: link.outputName, inputName: link.inputName}))
      }
    },
    dropThing (thing) {
      this.$socket.emit('updateThingsPosition', JSON.stringify({trackName: this.track.name, name: thing.name, x: thing.x, y: thing.y}))
    },
    updateSetting (setting) {
      this.$socket.emit('updateSetting', JSON.stringify({trackName: this.track.name, thingName: setting.blockName, name: setting.name, value: setting.value}))
    },
    clickSetting (setting) {
      if (this.$store.state.midi) {
        this.$store.state.midi = false
        this.$socket.emit('attachToMIDI', JSON.stringify({trackName: this.track.name, thingName: setting.blockName, name: setting.name}))
      }
    }
  },
  components: {
    BlocksContainer,
    AddThingDropdown
  }
}
</script>

<style scoped>
  #track-screen {
    height: 100%;
  }
</style>
