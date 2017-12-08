import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy, {
  defaultIconPack: 'fa'
})

import Checkin from './components/checkin/index.vue'
Vue.component('checkin', Checkin)

new Vue({
  el: "#app"
})
