<template>
  <div :style="{left: (dragging ? tx : x) + 'px', top: (dragging ? ty : y) + 'px'}" class="block">
    <div class="box">
      <div class="inputs">
        <div v-for="(input, index) in inputs" :key="input.ID">
          <span @mousedown="startLink($event, 'input', input, index)" @mouseup="stopLink($event, 'input', input, index)" class="tag">{{input.name}}</span>
        </div>
      </div>
      <div class="block-content">
        <h1 class="title">{{name}}</h1>
        <div v-for="(setting, index) in settings" @click="clickSetting($event, setting, index)" class="field is-grouped is-grouped-centered">
          <div v-if="setting.exposedName !== ''" class="control is-expanded">
            <input type="number" class="input is-small" :value="setting.min" @input="updateSetting($event, 'min', setting, index)">
          </div>
          <div v-if="setting.exposedName === ''" class="control is-expanded">
            <input type="number" class="input is-small" min="0" max="127" :value="setting.value" @input="updateSetting($event, 'value', setting, index)">
          </div>
          <div v-if="setting.exposedName !== ''" class="control is-expanded">
            <input type="number" class="input is-small" :value="setting.max" @input="updateSetting($event, 'max', setting, index)">
          </div>
          <button v-if="setting.exposedName === ''" class="button is-small" @click="exposeSetting(setting)">Expose</button>
          <button v-else class="button is-small" @click="hideSetting(setting)">Hide</button>
        </div>
      </div>
      <div class="outputs">
        <div v-for="(output, index) in outputs" :key="output.ID">
          <span @mousedown="startLink($event, 'output', output, index)" @mouseup="stopLink($event, 'output', output, index)" class="tag">{{output.name}}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Block',
  props: {
    x: Number,
    y: Number,
    name: String,
    settings: Array,
    inputs: Array,
    outputs: Array
  },
  data () {
    return {
      dragging: false,
      linking: false,
      lastMouseX: 0,
      lastMouseY: 0,
      tx: 0,
      ty: 0,
      height: 0,
      width: 0
    }
  },
  mounted () {
    this.height = this.$el.offsetHeight
    this.width = this.$el.offsetWidth

    this.$emit('init', this.height, this.width)

    document.documentElement.addEventListener('mousemove', this.mouseMove, true)
    document.documentElement.addEventListener('mousedown', this.mouseDown, true)
    document.documentElement.addEventListener('mouseup', this.mouseUp, true)
  },
  beforeDestroy () {
    document.documentElement.removeEventListener('mousemove', this.mouseMove, true)
    document.documentElement.removeEventListener('mousedown', this.mouseDown, true)
    document.documentElement.removeEventListener('mouseup', this.mouseUp, true)
  },
  methods: {
    mouseDown (e) {
      if (this.$el.contains(e.target)) {
        if (!e.target.classList.contains('tag')) {
          this.lastMouseX = e.pageX
          this.lastMouseY = e.pageY

          this.tx = this.x
          this.ty = this.y

          this.dragging = true
        }
      }
    },
    mouseUp (e) {
      if (this.dragging) {
        this.$emit('drop-block', this.tx, this.ty, this.height, this.width)

        this.dragging = false
      }

      if (this.linking && !e.target.classList.contains('tag')) {
        this.$emit('stop-link', e, '', null, -1)
      }

      this.linking = false
    },
    mouseMove (e) {
      if (this.dragging) {
        let lastMouseX = this.lastMouseX
        let lastMouseY = this.lastMouseY
        this.lastMouseX = e.pageX
        this.lastMouseY = e.pageY

        this.tx += this.lastMouseX - lastMouseX
        this.ty += this.lastMouseY - lastMouseY

        this.$emit('mousemove', this.tx, this.ty, this.height, this.width)
      } else if (this.linking) {
        this.$emit('move-link', e)
      }
    },
    startLink (e, type, elem, index) {
      this.linking = true
      this.$emit('start-link', e, type, elem, index)
    },
    stopLink (e, type, elem, index) {
      this.linking = false
      this.$emit('stop-link', e, type, elem, index)
    },
    updateSetting(e, kind, setting, index) {
      setting[kind] = Number(e.target.value)
      this.$emit('update-setting', setting)
    },
    clickSetting(e, setting, index) {
      this.$emit('click-setting', setting)
    },
    exposeSetting(setting) {
      this.$emit('expose-setting', setting)
    },
    hideSetting(setting) {
      this.$emit('hide-setting', setting)
    }
  }
}
</script>

<style scoped>
.block {
  position: absolute;
}

.box {
  margin-bottom: 0;
  display: flex;
  padding: 0;
}

.box .block-content {
  padding: 0 .25rem;
  text-align: center;
}

.block-content input {
  max-width: 55px;
}

.inputs {
  transform: translateX(-50%);
  display: -ms-flexbox;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.outputs {
  transform: translateX(50%);
  display: -ms-flexbox;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.inputs > div, .outputs > div {
  margin-top: 1px;
  margin-bottom: 1px;
}

.block {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.circle:hover {
  background: linear-gradient(to bottom, #ff4040, #f00);
}
</style>
