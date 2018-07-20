<template>
  <svg class="block-link">
    <path :d="d"></path>
  </svg>
</template>

<script>
export default {
  name: 'BlockLink',
  props: {
    p0: Object,
    p1: Object,
  },
  computed: {
    d () {
      let AX = this.p0.x
      let AY = this.p0.y

      let FX = this.p1.x
      let FY = this.p1.y

      let BX = Math.abs(FX - AX) * 0.05 + AX
      let BY = AY

      let CX = AX + Math.abs(FX - AX) * 0.33
      let CY = AY

      let DX = FX - Math.abs(FX - AX) * 0.33
      let DY = FY

      let EX = -1 * Math.abs(FX - AX) * 0.05 + FX
      let EY = FY

      return 'M' + AX + ',' + AY + ' L' + BX + ',' + BY + ' C' + CX + ',' + CY + ' ' + DX + ',' + DY + ' ' + EX + ',' + EY + ' L' + FX + ',' + FY
    }
  },
  data () {
    return {
    }
  }
}
</script>

<style scoped>
.block-link {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  fill: none;
  stroke-width: 3;
  stroke: black;
}
</style>
