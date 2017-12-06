import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

import Admin from './components/admin/index.vue'
Vue.component('admin', Admin)

new Vue({
  el: "#app"
})
