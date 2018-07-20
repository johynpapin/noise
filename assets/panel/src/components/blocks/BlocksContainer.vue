<template>
  <div class="blocks-container">
    <block-link v-if="tmpLink !== null" :p0="tmpLink.p0" :p1="tmpLink.p1"></block-link>
    <block-link v-for="(l, i) in links" :p0="l.p0" :p1="l.p1"></block-link>
    <block v-for="(b, i) in blocks" :x="b.x" :y="b.y" :name="b.name" :inputs="b.inputs" :outputs="b.outputs" @mousemove="(...args) => mouseMove(b, i, ...args)" @init="(...args) => init(b, i, ...args)" @move-link="(...args) => moveLink(b, i, ...args)" @start-link="(...args) => startLink(b, i, ...args)" @stop-link="(...args) => stopLink(b, i, ...args)"></block>
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
    blocks: Array
  },
  data () {
    return {
      links: {
      },
      tmpLink: null
    }
  },
  methods: {
    mouseMove (block, blockIndex, diffX, diffY) {
      this.blocks[blockIndex].x += diffX
      this.blocks[blockIndex].y += diffY
      this.updateLinks(block, blockIndex)
    },
    init (block, blockIndex, height, width) {
      this.$set(this.blocks[blockIndex], 'height', height)
      this.$set(this.blocks[blockIndex], 'width', width)
      this.updateLinks(block, blockIndex)
    },
    startLink (block, blockIndex, e, type, index, name) {
      if (type === 'output') {
        this.tmpLink = {
          block: block,
          type: type,
          index: index,
          name: name,
          p0: computeCirclePosition(block.x, block.y, block.height, block.width, index, type, 14, 10, block[type + 's'].length),
          p1: {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop}
        }
      } else {
        this.tmpLink = {
          block: block,
          type: type,
          index: index,
          name: name,
          p0: {x: e.pageX - this.$el.offsetLeft, y: e.pageY - this.$el.offsetTop},
          p1: computeCirclePosition(block.x, block.y, block.height, block.width, index, type, 14, 10, block[type + 's'].length)
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
    stopLink (block, blockIndex, e, type, index, name) {
      let startBlock = this.tmpLink.block
      let startType = this.tmpLink.type
      let startName = this.tmpLink.name
      let startIndex = this.tmpLink.index

      this.tmpLink = null
      if (index !== -1 && type !== startType) {
        if (startType === 'input') {
          [startBlock, block] = [block, startBlock];
          [startIndex, index] = [index, startIndex];
          [startName, name] = [name, startName];
        }

        if (!block.inputs[index].linkedTo) {
          this.$emit('new-link', {inputBlockName: block.name, inputName: name, outputBlockName: startBlock.name, outputName: startName})
        }
      }
    },
    updateLinks (block, blockIndex) {
      let blocksMap = {}
      for (let block of this.blocks) {
        blocksMap[block.name] = block
      }

      for (let outputIndex in block.outputs) {
        let linkName = block.name + block.outputs[outputIndex].name
        if (this.links[linkName]) {
          this.links[linkName].p0 = computeCirclePosition(block.x, block.y, block.height, block.width, outputIndex, 'output', 14, 10, block.outputs.length)
        }
      }

      for (let inputIndex in block.inputs) {
        let input = block.inputs[inputIndex]
        if (input.linkedTo) {
          let linkName = input.linkedTo.blockName + input.linkedTo.outputName

          if (!this.links[linkName]) {
            let outputBlock = blocksMap[input.linkedTo.blockName]
            let outputIndex = outputBlock.outputs.findIndex((e) => e.name === input.linkedTo.outputName)

            this.$set(this.links, linkName, {
              p1: {x: 0, y: 0},
              p0: computeCirclePosition(outputBlock.x, outputBlock.y, outputBlock.height, outputBlock.width, outputIndex, 'output', 14, 10, outputBlock.outputs.length)
            })
          }

          this.links[linkName].p1 = computeCirclePosition(block.x, block.y, block.height, block.width, inputIndex, 'input', 14, 10, block.inputs.length)
        }
      }
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