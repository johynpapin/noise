<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
@import "../node_modules/bulma/bulma.sass";

html, body, #app {
  height: 100%;
  width: 100%;
  overflow: hidden;
  background: #f5f5f5;
}
</style>
