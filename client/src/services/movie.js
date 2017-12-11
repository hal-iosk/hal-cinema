import axios from 'axios'
import CookieDoc from '../lib/CookieDoc'
import moment from 'moment'

const  MovieHttp = {

  // 全権取得
  GetMovies() {
    return axios.get(`/api/movie`)
  },
    //公開中映画
  GetOnAirMovies(date){
    return axios.get(`/api/movie?status=0&day=${date}`)
  },
    // 公開予定映画
  GetComingsoonMovies(){
    return axios.get(`/api/movie?status=1`)
  }
}

export default MovieHttp;