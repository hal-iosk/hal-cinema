import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

// components
import Admin from './components/admin/index.vue'
import MovieTableEdit from './components/admin/movieTableEdit.vue'
import Schedule from './components/admin/schedule.vue'

import NotFound from './components/notFound/index.vue'

// routing
const routes = [
  { path: "/admin", component: Admin },
  { path: "/admin/movieedit/:id", component: MovieTableEdit },
  { path: "/admin/schedule/:id", component: Schedule },
  { path: "*", component: NotFound },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
  router
}).$mount("#app")
