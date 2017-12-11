import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

import Movie from './components/movie/index.vue'

Vue.component("movie", Movie)

new Vue({
  el: "#app"
})
