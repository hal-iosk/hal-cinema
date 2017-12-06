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

}

export default new HttpUtils;
