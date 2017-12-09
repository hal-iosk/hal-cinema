import Vue from 'vue'

import Movie from './components/movie/index.vue'

Vue.component("movie", Movie)

new Vue({
  el: "#app"
})
