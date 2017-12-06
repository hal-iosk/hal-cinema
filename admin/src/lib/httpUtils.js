import axios from 'axios'

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
    var params = new URLSearchParams();
    params.append('movie_name', title);
    params.append('details', detail);
    params.append('image_path', thumbnail);
    params.append('start_date', start);
    params.append('end_date', end);
    params.append('watch_time', watch_time);

    return axios.put(`/api/admin/movie/${id}`, params)
  }

}

export default new HttpUtils;
