import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'
import Vuex from 'vuex'

Vue.use(Vuex)
Vue.use(VueRouter)
Vue.use(Buefy)

import Admin from './components/admin/index.vue'
import NotFound from './components/notFound/index.vue'

const routes = [
  { path: "/admin", component: Admin }
]

const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
