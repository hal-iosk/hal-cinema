import axios from 'axios'
import CookieDoc from '../lib/CookieDoc'

const ReserveHttp = {

  GetPrices() {
    return axios.get(`/api/price`)
  },

  PostReserve() {

  }

}

export default ReserveHttp;