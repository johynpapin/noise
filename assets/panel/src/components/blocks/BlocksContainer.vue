<template>
  <div class="blocks-container">
    <block-link v-if="tmpLink !== null" :p0="tmpLink.p0" :p1="tmpLink.p1"></block-link>
    <block-link v-for="(l, i) in links" :p0="l.p0" :p1="l.p1"></block-link>
    <block v-for="(b, i) in blocks" :key="b.ID" :x="b.x" :y="b.y" :name="b.name" :settings="b.settings" :inputs="b.inputs" :outputs="b.outputs" @expose-setting="exposeSetting(b, i, $event)" @hide-setting="hideSetting(b, i, $event)" @drop-block="(...args) => dropBlock(b, i, ...args)" @drag-start="dragStart(b, i)" @click-setting="(...args) => clickSetting(b, i, ...args)" @update-setting="(...args) => updateSetting(b, i, ...args)" @mousemove="(...args) => mouseMove(b, i, ...args)" @init="(...args) => init(b, i, ...args)" @move-link="(...args) => moveLink(b, i, ...args)" @start-link="(...args) => startLink(b, i, ...args)" @stop-link="(...args) => stopLink(b, i, ...args)"></block>
  </div>
</template>

<script>
import Block from './Block.vue'
import BlockLink from './BlockLink.vue'

function computeCirclePosition(blockX, blockY, blockHeight, blockWidth, x, circleType, circleDiameter, circleSpacing, numberOfCircles) {
  return {
    y: blockY + ((blockHeight - (numberOfCircles * (circleDiameter + circleSpacing) - circleSpacing)) / 2) + (circleDiameter + circleSpacing) * x + (circleDiameter / 2),
    x: blockX + (circleType === 'input' ? 0 : blockWidth)
  }
}

export default {
  name: 'BlocksContainer',
  props: {
    rawBlocks: Array
  },
  computed: {
    blocks () {
      let blocks = []

      for (let index in this.rawBlocks) {
        let block = this.rawBlocks[index]

        if (this.blocksMeta[block.ID]) {
          block = {...block, ...this.blocksMeta[block.ID]}
        }

        blocks.push(block)
      }

      return blocks
    },
    links () {
      let links = []

      if (!this.blocks) {
        return links
      }

      let blocksMap = {}

      for (let block of this.blocks) {
        blocksMap[block.ID] = block
      }

      for (let block of this.blocks) {
        if (!block.inputs) {
          continue
        }

        for (let inputIndex in block.inputs) {
          let input = block.inputs[inputIndex]

          if (input.linkedTo !== null) {
            let outputBlock = blocksMap[input.linkedTo.blockID]
            let outputIndex = 0

            for (let index in outputBlock.outputs) {
              if (outputBlock.outputs[index].ID === input.linkedTo.outputID) {
                outputIndex = index
              }
            }

            links.push({
              p0: computeCirclePosition(outputBlock.x, outputBlock.y, outputBlock.height, outputBlock.width, outputIndex, 'output', 24, 1, outputBlock.outputs.length),
              p1: computeCirclePosition(block.x, block.y, block.height, block.width, inputIndex, 'input', 14, 10, block.inputs.length)
            })
          }
        }
      }

      return links
    }
  },
  data () {
    return {
      tmpLink: null,
      blocksMeta: {}
    }
  },
  methods: {
    mouseMove (block, blockIndex, x, y) {
      let blockMeta = this.blocksMeta[block.ID]

      blockMeta.x = x
      blockMeta.y = y

      block.x = x
      block.y = y

      this.$set(this.blocksMeta, block.ID, blockMeta)
    },
    init (block, blockIndex, height, width) {
      this.$set(this.blocksMeta, block.ID, {width: width, height: height})

      block.width = width
      block.height = height
    },
    dropBlock (block, blockIndex, x, y) {
      block.x = x
      block.y = y

      this.$emit('drop-block', block)

      this.$set(this.blocksMeta[block.ID], 'dragging', false)
    },
    dragStart(block, blockIndex) {
      this.$set(this.blocksMeta[block.ID], 'dragging', true)
    },
    startLink (block, blockIndex, e, type, elem, index) {
      if (type === 'output') {
        this.tmpLink = {
          block: block,
          type: type,
          elem: elem,
          p0: computeCirclePosition(block.x, block.y, block.height, block.width, index, type, 24, 1, block[type + 's'].length),
          p1: {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop}
        }
      } else {
        this.tmpLink = {
          block: block,
          type: type,
          elem: elem,
          p0: {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop},
          p1: computeCirclePosition(block.x, block.y, block.height, block.width, index, type, 24, 1, block[type + 's'].length)
        }
      }
    },
    moveLink (block, blockIndex, e) {
      if (this.tmpLink.type === 'output') {
        this.tmpLink.p1 = {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop}
      } else {
        this.tmpLink.p0 = {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop}
      }
    },
    stopLink (block, blockIndex, e, type, elem, index) {
      if (index !== -1) {
        let input = elem
        let output = this.tmpLink.elem

        if (type === 'output') {
          [input, output] = [output, input]
        }

        if (input.linkedTo && input.linkedTo.ID === output.ID) {
          this.$emit('destroy-link', {input: input, output: output})
        } else if (!input.linkedTo && type !== this.tmpLink.type) {
          this.$emit('new-link', {input: input, output: output})
        }
      }

      this.tmpLink = null
    },
    updateSetting (block, blockIndex, setting) {
      this.$emit('update-setting', {thing: block, setting: setting})
    },
    clickSetting (block, blockIndex, setting) {
      this.$emit('click-setting', {thing: block, setting: setting})
    },
    exposeSetting (block, blockIndex, setting) {
      this.$emit('expose-setting', {thing: block, setting: setting})
    },
    hideSetting (block, blockIndex, setting) {
      this.$emit('hide-setting', {thing: block, setting: setting})
    }
  },
  components: {
    Block,
    BlockLink
  }
}
</script>

<style scoped>
.blocks-container {
  position: relative;
  height: 100%;
  width: 100%;
}
</style>
