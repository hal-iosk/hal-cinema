import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    movie_title: "",
    watch_time: 0
  },
  mutations: {
    set_movie_title (state, title) {
      state.movie_title = title;
    },
    set_watch_time (state, time) {
      state.watch_time = time;
    }
  }
})

export default store;
