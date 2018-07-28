<template>
  <div class="sidebar">
    <tracks></tracks>
    <instruments></instruments>
  </div>
</template>

<script>
import Tracks from './Tracks.vue'
import Instruments from './Instruments.vue'

export default {
  name: 'Sidebar',
  data () {
    return {
    }
  },
  components: {
    Tracks,
    Instruments
  }
}
</script>

<style scoped>
.sidebar {
  height: 100%;
  width: 100%;

  background: #fff;
  box-shadow: 0 2px 3px hsla(0,0%,7%,.1),0 0 0 1px hsla(0,0%,7%,.1);
}
</style>
