import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex);

const store = new Vuex.Store({
  plugins: [createLogger()],
  state: {
    movie_title: "",
    watch_time: 0,
    coin: 0
  },
  mutations: {
    set_movie_title (state, title) {
      state.movie_title = title;
    },
    set_watch_time (state, time) {
      state.watch_time = time;
    },
    coinUpdate(state, coin) {
      state.coin = coin;
    }
  }
})

export default store;
