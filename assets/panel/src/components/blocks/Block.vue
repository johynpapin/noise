<template>
  <div :style="{left: x + 'px', top: y + 'px'}" class="block">
    <div class="inputs">
      <div v-for="(input, index) in inputs" @mousedown="startLink($event, 'input', index, input.name)" @mouseup="stopLink($event, 'input', index, input.name)" class="circle"></div>
    </div>
    <div class="box">
      <h1 class="title">{{name}}</h1>
    </div>
    <div class="outputs">
      <div v-for="(output, index) in outputs" @mousedown="startLink($event, 'output', index, output.name)" @mouseup="stopLink($event, 'output', index, output.name)" class="circle"></div>
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
    inputs: Array,
    outputs: Array
  },
  data () {
    return {
      dragging: false,
      linking: false,
      lastMouseX: 0,
      lastMouseY: 0,
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
        if (!e.target.classList.contains('circle')) {
          this.lastMouseX = e.pageX
          this.lastMouseY = e.pageY

          this.dragging = true
        }
      }
    },
    mouseUp (e) {
      this.dragging = false

      if (this.linking && !e.target.classList.contains('circle')) {
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

        this.$emit('mousemove', this.lastMouseX - lastMouseX, this.lastMouseY - lastMouseY, this.height, this.width)
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
}

.inputs {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  left: -7px;
}

.outputs {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  right: -7px;
}

.circle {
  margin-top: 10px;
  margin-bottom: 10px;

  position: relative;
  width: 14px;
  height: 14px;
  border-radius: 7px;
  background: linear-gradient(to bottom, #40b3ff, #09f);
}

.circle:hover {
  background: linear-gradient(to bottom, #ff4040, #f00);
}
</style>
