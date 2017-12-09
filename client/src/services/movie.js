import axios from 'axios'
import CookieDoc from '../lib/CookieDoc'
import moment from 'moment'

const  MovieHttp = {

  GetMovies() {
    return axios.get(`/api/movie`)
  }

}

export default MovieHttp;