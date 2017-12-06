import Vue from 'vue'
import Vuex from 'vuex'
import mockMovies from '../mock/movies'

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    movies: mockMovies
  },
  mutations: {
    hoge (state) {
      state.count++
    }
  }
})

export default store;
