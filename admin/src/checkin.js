import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy, {
  defaultIconPack: 'fa'
})

import Checkin from './components/checkin/index.vue'
import Login from './components/checkin/login.vue'

// routing
const routes = [
  { path: "/checkin", component: Checkin },
  { path: "/checkin/login", component: Login }
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")