<template>
  <div :style="{left: (dragging ? tx : x) + 'px', top: (dragging ? ty : y) + 'px'}" class="block">
    <div class="box">
      <div class="inputs">
        <div v-for="(input, index) in inputs" :key="name + 'i' + input.name">
          <span @mousedown="startLink($event, 'input', index, input.name)" @mouseup="stopLink($event, 'input', index, input.name)" class="tag">{{input.name}}</span>
        </div>
      </div>
      <div class="block-content">
        <h1 class="title">{{name}}</h1>
        <div v-for="(setting, index) in settings" @click="clickSetting($event, setting, index)" class="field">
          <label class="label">{{setting.name}}</label>
          <div class="control">
            <input class="input is-rounded" type="number" :value="setting.value" @input="updateSetting($event, setting, index)">
          </div>
        </div>
      </div>
      <div class="outputs">
        <div v-for="(output, index) in outputs" :key="name + 'o' + output.name">
          <span @mousedown="startLink($event, 'output', index, output.name)" @mouseup="stopLink($event, 'output', index, output.name)" class="tag">{{output.name}}</span>
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
        this.$emit('stop-link', e, '', -1)
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
    startLink (e, type, index, name) {
      this.linking = true
      this.$emit('start-link', e, type, index, name)
    },
    stopLink (e, type, index, name) {
      this.linking = false
      this.$emit('stop-link', e, type, index, name)
    },
    updateSetting(e, setting, index) {
      this.$emit('update-setting', setting.name, Number(e.target.value))
    },
    clickSetting(e, setting, index) {
      this.$emit('click-setting', setting.name)
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
