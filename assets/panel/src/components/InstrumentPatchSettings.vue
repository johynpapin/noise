<template>
  <div class="instrument-patch-settings">
    <div class="columns is-multiline">
      <div class="column is-narrow" :key="setting.ID" v-for="setting in patch.settings">
        <div class="field" :class="{'has-background-primary': setting.attached, 'has-background-warning': MIDIEditMode && !setting.attached}" @click="attachToMIDI($event, setting)">
          <label class="label">{{setting.exposedName}}</label>
          <div class="control">
            <input class="input" type="number" :value="setting.value" @input="updateSetting($event, setting)" :min="setting.min" :max="setting.max" :step="0.001">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'InstrumentPatchSettings',
  props: {
    patch: Object
  },
  data () {
    return {
    }
  },
  computed: mapState({
    MIDIEditMode: state => state.MIDIEditMode
  }),
  methods: {
    attachToMIDI (e, setting) {
      this.$socket.emit('command', JSON.stringify({command: 'attachToMIDI', instrumentPatchID: this.patch.ID, settingID: setting.ID}))
    },
    updateSetting (e, setting) {
      setting.value = Number(e.target.value)

      this.$socket.emit('command', JSON.stringify({command: 'updateExposedSetting', instrumentPatchID: this.patch.ID, settingID: setting.ID, value: setting.value}))
    }
  }
}
</script>
