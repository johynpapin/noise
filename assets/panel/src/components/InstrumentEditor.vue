<template>
  <div id="instrument-editor">
    <add-thing-dropdown :instrument="instrument"></add-thing-dropdown>
    <blocks-container @drop-block="dropThing" @update-setting="updateSetting" @expose-setting="exposeSetting" @hide-setting="hideSetting" @destroy-link="destroyLink" @new-link="newLink" :raw-blocks="blocks"></blocks-container>
  </div>
</template>

<script>
import BlocksContainer from './blocks/BlocksContainer.vue'
import AddThingDropdown from './AddThingDropdown.vue'

export default {
  name: 'InstrumentEditor',
  props: {
    instrument: Object
  },
  computed: {
    blocks () {
      let blocks = [
        {
          ID: -2,
          name: "Olala",
          outputs: [this.instrument.entryPoint],
          x: 40,
          y: 100
        }, {
          ID: -3,
          name: "Olali",
          inputs: [this.instrument.exitPoint],
          x: 1000,
          y: 100
        }
      ]

      if (this.instrument.exitPoint.linkedTo) {
        blocks[1].inputs[0].linkedTo.blockID = blocks[1].inputs[0].linkedTo.thingID
      }

      if (!this.instrument.things) {
        return blocks
      }

      for (let block of this.instrument.things) {
        for (let inputIndex in block.inputs) {
          if (block.inputs[inputIndex].linkedTo) {
            block.inputs[inputIndex].linkedTo.blockID = block.inputs[inputIndex].linkedTo.thingID
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
      if ((link.input.ID < 0 && link.output.ID < 0) || link.input.ID === link.output.ID) {
        return
      }

      if (link.output.ID < 0) {
        this.$socket.emit('command', JSON.stringify({command: 'attachToInstrumentEntryPoint', instrumentID: this.instrument.ID, inputID: link.input.ID}))
      } else if (link.input.ID < 0) {
        this.$socket.emit('command', JSON.stringify({command: 'attachToInstrumentExitPoint', instrumentID: this.instrument.ID, outputID: link.output.ID}))
      } else {
        this.$socket.emit('command', JSON.stringify({command: 'attachToThing', inputID: link.input.ID, outputID: link.output.ID}))
      }
    },
    destroyLink (link) {
      if ((link.input.ID < 0 && link.output.ID < 0) || link.input.ID === link.output.ID) {
        return
      }

      if (link.output.ID < 0) {
        this.$socket.emit('command', JSON.stringify({command: 'detachFromInstrumentEntryPoint', instrumentID: this.instrument.ID, inputID: link.input.ID}))
      } else if (link.input.ID < 0) {
        this.$socket.emit('command', JSON.stringify({command: 'detachFromInstrumentExitPoint', instrumentID: this.instrument.ID, outputID: link.output.ID}))
      } else {
        this.$socket.emit('command', JSON.stringify({command: 'detachFromThing', inputID: link.input.ID, outputID: link.output.ID}))
      }
    },
    dropThing (thing) {
      this.$socket.emit('command', JSON.stringify({command: 'updateThingPosition', thingID: thing.ID, x: thing.x, y: thing.y}))
    },
    updateSetting ({thing, setting}) {
      this.$socket.emit('command', JSON.stringify({command: 'updateSetting', settingID: setting.ID, name: setting.exposedName, value: setting.value, min: setting.min, max: setting.max}))
    },
    exposeSetting ({thing, setting}) {
      this.$socket.emit('command', JSON.stringify({command: 'exposeSetting', settingID: setting.ID, instrumentID: this.instrument.ID, name: thing.name + '_' + setting.name}))
    },
    hideSetting ({thing, setting}) {
      this.$socket.emit('command', JSON.stringify({command: 'hideSetting', settingID: setting.ID, instrumentID: this.instrument.ID}))
    },
  },
  components: {
    BlocksContainer,
    AddThingDropdown
  }
}
</script>

<style scoped>
#instrument-editor {
  height: 100%;
}
</style>
