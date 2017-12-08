import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

import Reserve from './components/reserve'

// routing
const routes = [
  { path: "/", component: Reserve },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
