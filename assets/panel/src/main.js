import Vue from 'vue'
import App from './App'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faMusic, faPlus } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import VueSocketio from 'vue-socket.io'

import router from './router'
import store from './store'

Vue.config.productionTip = false

library.add(faMusic, faPlus)

Vue.component('fa', FontAwesomeIcon)

Vue.use(VueSocketio, 'http://localhost:4242', store)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
