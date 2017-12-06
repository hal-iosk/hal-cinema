import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

import Login from './components/login/index.vue'
Vue.component('login', Login)

new Vue({
  el: "#app"
})
