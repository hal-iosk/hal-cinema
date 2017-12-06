import axios from 'axios'
import CookieDoc from './CookieDoc'
import moment from 'moment'

class HttpUtils {

  Login(email, password) {
    var params = new URLSearchParams();
    params.append('email', email);
    params.append('password', password);

    return axios.post("/api/admin/signin", params)
  }

  GetMovies() {
    return axios.get("/api/movie")
  }

  GetMovieDetail(id) {
    return axios.get(`/api/movie/${id}`)
  }

  PutMovieDetail(id, title, detail, thumbnail, start, end, watch_time) {
    const token = CookieDoc.getItem("halCinemaAdmin")

    var params = new URLSearchParams();
    params.append('movie_name', title);
    params.append('details', detail);
    params.append('image_path', thumbnail);
    params.append('start_date', moment(start).format("YYYY/MM/DD"));
    params.append('end_date', moment(end).format("YYYY/MM/DD"));
    params.append('watch_time', watch_time);

    return axios.put(`/api/admin/movie/${id}?token=${token}`, params)
  }

  DeleteMovieDetail(id) {
    const token = CookieDoc.getItem("halCinemaAdmin")
    return axios.delete(`/api/admin/movie/${id}?token=${token}`)
  }

}

export default new HttpUtils;
