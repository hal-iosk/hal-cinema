import Vue from 'vue'

import Reserve from './components/reserve/index.vue'
import Movie from './components/movie/index.vue'

Vue.component("movie", Movie)

new Vue({
  el: "#app"
})
