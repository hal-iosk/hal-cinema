import axios from 'axios'
import CookieDoc from '../lib/CookieDoc'

const ReserveHttp = {

  GetPrices() {
    return axios.get(`/api/price`)
  },

  PostReserve(movie, seats) {
    const token = CookieDoc.getItem("halCinemaUser")

    const params = seats.map((seat) => {
      return {
        schedule_id: movie.scheduleId,
        seat_id: seat.seatId,
        price_id: seat.priceId
      }
    });

    return axios.post(`/api/reserve?token=${token}`, {
      reserves: params
    })
  }

}

export default ReserveHttp;