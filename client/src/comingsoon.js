import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

import Comingsoon from './components/comingsoon/index.vue'

Vue.component("comingsoon", Comingsoon)

new Vue({
    el: "#app"
})
