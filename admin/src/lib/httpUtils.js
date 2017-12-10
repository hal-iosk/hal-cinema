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

  UserLogin(email, password) {
    var params = new URLSearchParams();
    params.append('email', email);
    params.append('password', password);

    return axios.post("/api/signin", params)
  }

  GetMovies() {
    return axios.get("/api/movie")
  }

  GetMovieDetail(id) {
    return axios.get(`/api/movie/${id}`)
  }

  CreateMovie(title, detail, thumbnail, start, end, watch_time) {
    const token = CookieDoc.getItem("halCinemaAdmin")

    var params = new URLSearchParams();
    params.append('movie_name', title);
    params.append('details', detail);
    params.append('image_path', thumbnail);
    params.append('start_date', moment(start).format("YYYY/MM/DD"));
    params.append('end_date', moment(end).format("YYYY/MM/DD"));
    params.append('watch_time', watch_time);

    return axios.post(`/api/admin/movie?token=${token}`, params)
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

  GetSchedule(id) {
    return axios.get(`/api/schedule?movie_id=${id}`)
  }

  CreateSchedule(movie_id, theater_number, start_time) {
    const token = CookieDoc.getItem("halCinemaAdmin")

    var params = new URLSearchParams();
    params.append('movie_id', movie_id);
    params.append('theater_number', theater_number);
    params.append('start_time', moment(start_time).format("YYYY/MM/MM hh:mm:ss"));

    return axios.post(`/api/admin/schedule?token=${token}`, params)
  }

  PutSchedule(schedule) {
    const token = CookieDoc.getItem("halCinemaAdmin")

    var params = new URLSearchParams();
    params.append('movie_id', schedule.movie_id)
    params.append('release', schedule.release)
    params.append('start_time', moment(schedule.start_time).format("YYYY/MM/MM hh:mm:ss"))
    params.append('theater_number', schedule.theater_number)

    return axios.put(`/api/admin/schedule/${schedule.ID}?token=${token}`, params)
  }

  DeleteSchedule(id) {
    const token = CookieDoc.getItem("halCinemaAdmin")
    return axios.delete(`/api/admin/schedule/${id}?token=${token}`)
  }

  PostCheckin() {
    const token = CookieDoc.getItem("halCinemaUser")
    return axios.post(`/api/checkin?token=${token}`)
  }

  PostPopcorn() {
    const token = CookieDoc.getItem("halCinemaUser")
    return axios.post(`/api/popcorn?token=${token}`)
  }

}

export default new HttpUtils;
