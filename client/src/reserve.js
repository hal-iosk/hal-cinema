import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

import Reserve from './components/reserve/index.vue'
import TicketType from './components/reserve/ticketType.vue'
import ReserveComplete from './components/reserve/reserveConfirm.vue'

// routing
const routes = [
  { path: "/reserve", component: Reserve },
  { path: "/reserve/ticket", component: TicketType },
  { path: "/reserve/complete", component: ReserveComplete },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
