import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

import TicketType from './components/reserve/ticketType.vue'

// routing
const routes = [
  { path: "/reserve", component: TicketType },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
