import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex);

const store = new Vuex.Store({
  plugins: [createLogger()],
  state: {
  },
  mutations: {
  }
})

export default store;
