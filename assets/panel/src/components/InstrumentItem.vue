<template>
  <a class="panel-block" :class="[{'is-active': selected}]">
    <span class="panel-icon">
      <fa icon="music"></fa>
    </span>
    {{name}}
  </a>
</template>

<script>
export default {
  name: 'InstrumentItem',
  props: {
    name: String,
    selected: Boolean
  },
  data () {
    return {
    }
  }
}
</script>
