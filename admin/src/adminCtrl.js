import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

// components
import Admin from './components/admin/index.vue'
import SalesManageEdit from './components/admin/salesManageEdit.vue'
import MovieTableEdit from './components/admin/movieTableEdit.vue'

import NotFound from './components/notFound/index.vue'

// routing
const routes = [
  { path: "/admin", component: Admin },
  { path: "/admin/salesedit", component: SalesManageEdit },
  { path: "/admin/movieedit", component: MovieTableEdit },
  { path: "*", component: NotFound },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
