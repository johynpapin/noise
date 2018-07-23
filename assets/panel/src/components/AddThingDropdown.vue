<template>
  <div class="add-thing-dropdown">
    <div class="dropdown is-hoverable">
      <div class="dropdown-trigger">
        <button class="button is-large is-primary is-inverted" aria-haspopup="true" aria-controls="add-thing-menu">
          <span class="icon">
            <fa icon="plus"></fa>
          </span>
          <span>Add a thing</span>
        </button>
        <div class="dropdown-menu" id="add-thing-menu" role="menu">
          <div class="dropdown-content">
            <a @click="openAddThingModal(thing)" v-for="thing in things" :key="thing" class="dropdown-item">
              {{thing}}
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="modal" :class="[{'is-active': thing !== null}]">
      <div class="modal-background" @click="closeAddThingModal"></div>
      <div class="modal-content">
        <div class="box">
          <h1 class="title" style="color: #000;">Add a <strong>{{thing}}</strong></h1>
          <form @submit="addThing">
            <div class="field">
              <label class="label">Thing name</label>
              <div class="control">
                <input class="input" type="text" v-model="thingName" placeholder="Thing name">
              </div>
            </div>
            <div class="control">
              <button class="button is-primary">Add thing</button>
            </div>
          </form>
        </div>
      </div>
      <button class="modal-close is-large" aria-label="close" @click="closeAddThingModal"></button>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'AddThingDropdown',
  props: {
    track: Object
  },
  data () {
    return {
      thing: null,
      thingName: null
    }
  },
  computed: mapState({
    things: state => state.things
  }),
  methods: {
    openAddThingModal (thing) {
      this.thing = thing
    },
    closeAddThingModal () {
      this.thing = null
      this.thingName = null
    },
    addThing () {
      this.$socket.emit('createThing', JSON.stringify({trackName: this.track.name, name: this.thingName, kind: this.thing}))
      this.thing = null
      this.thingName = null
    }
  }
}
</script>
